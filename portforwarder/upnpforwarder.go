package portforwarder

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func sendMappingRequest(interfaces map[string]bool, action string, req string) bool {
	client := &http.Client{}
	for addr := range interfaces {
		req, err := http.NewRequest("POST", addr, strings.NewReader(req))
		if err == nil {
			req.Header.Add("Content-type", `text/xml; charset="utf-8"`)
			req.Header.Add("SOAPAction", action)
			resp, _ := client.Do(req)
			if resp.StatusCode == http.StatusOK {
				return true
			}
		} else {
			fmt.Println(err)
		}
	}
	return false
}

/*
 0 - add mapping
 1 - remove mapping
*/
func (upnpForwarder UPNPForwarder) changeForwarding(protocol string, mode int, readyStream chan string) {
	protocol = strings.ToUpper(protocol)
	var upnpInformationStream chan UPNPInformation
	upnpInformationStream = make(chan UPNPInformation)
	fmt.Println(upnpForwarder.upnpInterfaces)
	localIp := GetOutboundIP()
	var mRes bool
	if len(upnpForwarder.upnpInterfaces) == 0 {
		upnpForwarder.ssdpReq(upnpInformationStream)
		for {
			select {
			case upnpInformation := <-upnpInformationStream:
				if upnpInformation.Key == "location" {
					locationS := upnpInformation.Value
					locationUrl, err := url.Parse(locationS)
					if err != nil {
						fmt.Println("Can't parse string")
						continue
					}
					fmt.Printf("Got from upnpInformationStream: location is %s!\n", locationS)
					resp, err := http.Get(upnpInformation.Value)
					if err != nil {
						fmt.Println("Error: can't get information from router")
						continue
					}
					defer resp.Body.Close()
					data, err := io.ReadAll(resp.Body)
					if err != nil {
						fmt.Println("Error: can't read information from router")
						continue
					}
					sdata := string(data)
					ParseUPNPInterfaces(sdata, locationUrl.Scheme+"://"+locationUrl.Host, &upnpForwarder.upnpInterfaces)
					if mode == 0 {
						mRes = sendMappingRequest(upnpForwarder.upnpInterfaces, addPortMappingUrl, fmt.Sprintf(addPortMappingRequestFormat,
							upnpForwarder.OuterPort,
							protocol,
							upnpForwarder.InnerPort,
							localIp))
					} else if mode == 1 {
						mRes = sendMappingRequest(upnpForwarder.upnpInterfaces, deletePortMappingUrl, fmt.Sprintf(deletePortMappingRequestFormat,
							protocol,
							upnpForwarder.OuterPort))
					}
					if mRes {
						readyStream <- "finished"
						return
					}
				}
			}
		}
	} else {
		fmt.Println(1212)
		if mode == 0 {
			mRes = sendMappingRequest(upnpForwarder.upnpInterfaces, addPortMappingUrl, fmt.Sprintf(addPortMappingRequestFormat,
				upnpForwarder.OuterPort,
				upnpForwarder.InnerPort,
				localIp))
		} else if mode == 1 {
			mRes = sendMappingRequest(upnpForwarder.upnpInterfaces, deletePortMappingUrl, fmt.Sprintf(deletePortMappingRequestFormat,
				upnpForwarder.OuterPort))
		}
		if mRes {
			readyStream <- "finished"
			return
		}
	}
}
func (upnpForwarder UPNPForwarder) ssdpReq(upnpInformationStream chan UPNPInformation) {
	laddr, err := net.ResolveUDPAddr(ssdpProto, ssdpBindAddr+":"+strconv.Itoa(upnpForwarder.SsdpPort))
	if err != nil {
		panic(err)
	}
	raddr, _ := net.ResolveUDPAddr(ssdpProto, ssdpAddr+":"+strconv.Itoa(ssdpServerPort))
	var readyStream chan string
	readyStream = make(chan string)
	go upnpForwarder.listenForSsdpResponse(readyStream, upnpInformationStream)
	_ = <-readyStream //ready

	conn, err := net.ListenUDP(ssdpProto, laddr)
	defer conn.Close()
	if err != nil {
		fmt.Printf("Cannot dial to this addr %s:%d\nError: %s\n", ssdpBindAddr, upnpForwarder.SsdpPort, err.Error())
		os.Exit(1)
	}
	fmt.Println(raddr)
	conn.WriteToUDP([]byte(ssdpSearchRequest), raddr)
}

func (upnpForwarder UPNPForwarder) listenForSsdpResponse(readyStream chan string, upnpInformationStream chan UPNPInformation) {
	addr := net.UDPAddr{
		Port: upnpForwarder.SsdpPort,
		IP:   GetOutboundIP(),
	}
	conn, err := net.ListenUDP(ssdpProto, &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var buf [2048]byte
	readyStream <- "ready"
	for {
		_, _, err = conn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}
		responseString := string(buf[:])
		params := ParseSSDPResponse(responseString)
		//fmt.Println(params);
		st, okSt := params["ST"]
		location, okLoc := params["LOCATION"]
		if okSt && okLoc && st == "upnp:rootdevice" {
			fmt.Printf("Root upnp device has location: %s\n", location)
			upnpInformationStream <- UPNPInformation{Key: "location", Value: location}
		}
	}
}

func NewUPNPForwarder(ssdpPort int, innerPort int, outerPort int) *UPNPForwarder {
	pf := new(UPNPForwarder)
	pf.SsdpPort = ssdpPort
	pf.upnpInterfaces = make(map[string]bool)
	pf.InnerPort = innerPort
	pf.OuterPort = outerPort
	return pf
}
