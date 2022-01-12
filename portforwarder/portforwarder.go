package portforwarder

func StartForwarding(protocol string, forwarder Portforwarder, readyStream chan string) {
	go forwarder.changeForwarding(protocol, 0, readyStream)
}
func StopForwarding(protocol string, forwarder Portforwarder, readyStream chan string) {
	go forwarder.changeForwarding(protocol, 1, readyStream)
}
