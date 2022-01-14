package networking

import (
	"fmt"
	"net"
	"strconv"
	"github.com/kk222mo/godist/config"
	"github.com/kk222mo/godist/netutils"
)

func handleClientTCP(conn net.Conn) {
	data := make([]byte, config.NODE_PACKET_SIZE)
	for {
		_, err := conn.Read(data)
		if err != nil {
			fmt.Println("Some bad information from socket")
		} else {

		}
	}

	conn.Close()
}

func ServeTCP() {
	listener, err := net.Listen("tcp", netutils.GetOutboundIP().String() + ":" + strconv.Itoa(config.DEFAULT_INTERNAL_PORT))
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
	listener, err := net.Listen("udp", netutils.GetOutboundIP().String() + ":" + strconv.Itoa(config.DEFAULT_INTERNAL_PORT))
	defer listener.Close()
	if err != nil {
		fmt.Println("Error: Can't start UDP listening:", err)
		return
	}
	
}
