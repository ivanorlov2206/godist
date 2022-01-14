package netutils

import (
  "net"
  "bytes"
  "io/ioutil"
  "net/http"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func GetExternalIP() (net.IP, error) {
  resp, err := http.Get("http://checkip.amazonaws.com")
  if err != nil {
    return net.IP{}, err
  }
  defer resp.Body.Close()
  buf, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return net.IP{}, err
  }
  return net.ParseIP(string(bytes.TrimSpace(buf))), nil
}
