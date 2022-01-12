package main

import (
	"fmt"
	"github.com/kk222mo/godist/portforwarder"
)

func main() {
	fmt.Println("Starting port forwarding")
	forwarder := portforwarder.NewUPNPForwarder(22222, 33334, 33334)
	readyStream := make(chan string)
	portforwarder.StartForwarding("tcp", forwarder, readyStream)
	<-readyStream
	for {

	}
}
