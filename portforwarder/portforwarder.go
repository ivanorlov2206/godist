package portforwarder

func StartForwarding(forwarder Portforwarder, readyStream chan string) {
	go forwarder.changeForwarding(0, readyStream)
}
func StopForwarding(forwarder Portforwarder, readyStream chan string) {
	go forwarder.changeForwarding(1, readyStream)
}
