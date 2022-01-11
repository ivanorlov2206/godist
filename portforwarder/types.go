package portforwarder

type Portforwarder interface {
	forward()
}

const (
	ssdpBindAddr   = "0.0.0.0"
	ssdpProto      = "udp"
	ssdpAddr       = "239.255.255.250"
	ssdpServerPort = "1900"

	ssdpSearchRequest = "M-SEARCH * HTTP/1.1\r\n" +
		"ST: upnp:rootdevice\r\n" +
		"MX: 3\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"HOST: 239.255.255.250:1900\r\n\r\n"
)

type UPNPForwarder struct {
	SsdpPort  string
	InnerPort string
	OuterPort string
}
