package portforwarder

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func (upnpForwarder UPNPForwarder) forward() {
	upnpForwarder.ssdpReq()
}
func (upnpForwarder UPNPForwarder) ssdpReq() {
	laddr, err := net.ResolveUDPAddr(ssdpProto, ssdpBindAddr+":"+upnpForwarder.SsdpPort)
	ssdpServerPortI, _ := strconv.Atoi(ssdpServerPort)
	raddr := net.UDPAddr{IP: net.ParseIP(ssdpAddr), Port: ssdpServerPortI}
	var commandStream chan string
	commandStream = make(chan string)
	go upnpForwarder.listenForSsdpResponse(commandStream)
	_ = <-commandStream

	conn, err := net.DialUDP(ssdpProto, laddr, &raddr)
	if err != nil {
		fmt.Printf("Cannot dial to this addr %s:%s\nError: %s\n", ssdpBindAddr, upnpForwarder.SsdpPort, err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(conn, ssdpSearchRequest)

	defer conn.Close()
}

func (upnpForwarder UPNPForwarder) listenForSsdpResponse(commandStream chan string) {
	ssdpPortI, _ := strconv.Atoi(upnpForwarder.SsdpPort)
	addr := net.UDPAddr{
		Port: ssdpPortI,
		IP:   GetOutboundIP(),
	}
	conn, err := net.ListenUDP(ssdpProto, &addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	var buf [2048]byte

	for {
		commandStream <- "ready"
		_, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			panic(err)
		}
		responseString := string(buf[:])
		responseLines := strings.Split(responseString, "\r\n")
		for i := 1; i < len(responseLines); i++ {
			fmt.Println(responseLines[i])
		}
	}
}

func NewUPNPForwarder(ssdpPort string, innerPort string, outerPort string) *UPNPForwarder {
	pf := new(UPNPForwarder)
	pf.SsdpPort = ssdpPort
	pf.InnerPort = innerPort
	pf.OuterPort = outerPort
	return pf
}
