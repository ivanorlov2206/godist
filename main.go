package main

import (
	"fmt"
	"github.com/kk222mo/godist/portforwarder"
)

func main() {
	fmt.Println("Fuck")
	forwarder := portforwarder.NewUPNPForwarder("33333", "33334", "33334")
	portforwarder.StartForwarding(forwarder)
	for {

	}
}
