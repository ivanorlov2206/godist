package networking

import (
	"fmt"
	"github.com/kk222mo/godist/config"
	"github.com/kk222mo/godist/netutils"
	"net"
	"strconv"
)

func handleClientTCP(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, config.NODE_PACKET_SIZE)
	for {
		_, err := conn.Read(data)
		if err != nil {
			fmt.Println("Some bad information from socket")
			continue
		}
		nodeMessage, err := ParseNodeMessage(data)
		if (err != nil) {
			continue
		}
		fmt.Println(nodeMessage)
	}

}

func ServeTCP() {
	listener, err := net.Listen("tcp", netutils.GetOutboundIP().String()+":"+strconv.Itoa(config.DEFAULT_INTERNAL_PORT))
	defer listener.Close()
	if err != nil {
		fmt.Println("Error: Can't start TCP listening:", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClientTCP(conn)
	}
}

func ServeUDP() {
	addr := net.UDPAddr{
		Port: config.DEFAULT_INTERNAL_PORT,
		IP: netutils.GetOutboundIP()
	}
	listener, err := net.ListenUDP("udp", addr)
	defer listener.Close()
	if err != nil {
		fmt.Println("Error: Can't start UDP listening:", err)
		return
	}
	for {
		rlen, remote, err := listener.
	}
}
