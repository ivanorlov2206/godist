package portforwarder

type Portforwarder interface {
	changeForwarding(mode int, readyStream chan string)
}

const (
	ssdpBindAddr   = "0.0.0.0"
	ssdpProto      = "udp"
	ssdpAddr       = "239.255.255.250"
	ssdpServerPort = 1900

	ssdpSearchRequest = "M-SEARCH * HTTP/1.1\r\n" +
		"ST: upnp:rootdevice\r\n" +
		"MX: 3\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"HOST: 239.255.255.250:1900\r\n\r\n"

	addPortMappingRequestFormat = `<?xml version="1.0" encoding="utf-8"?>
	<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
	      <s:Body>
	        <u:AddPortMapping xmlns:u="urn:schemas-upnp-org:service:WANIPConnection:1">
	          <NewRemoteHost></NewRemoteHost>
	          <NewExternalPort>%d</NewExternalPort>
	          <NewProtocol>UDP</NewProtocol>
	          <NewInternalPort>%d</NewInternalPort>
	          <NewInternalClient>%s</NewInternalClient>
	          <NewEnabled>1</NewEnabled>
	          <NewPortMappingDescription>Vanya's UPnP UDP Golang</NewPortMappingDescription>
	          <NewLeaseDuration>0</NewLeaseDuration>
	        </u:AddPortMapping>
	      </s:Body>
	    </s:Envelope>`
	addPortMappingUrl = `"urn:schemas-upnp-org:service:WANIPConnection:1#AddPortMapping"`

	deletePortMappingRequestFormat = `
	<?xml version="1.0" encoding="utf-8"?>
	<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
		<s:Body>
			<u:DeletePortMapping xmlns:u="urn:schemas-upnp-org:service:WANIPConnection:1">
				<NewRemoteHost></NewRemoteHost>
				<NewProtocol>UDP</NewProtocol>
				<NewExternalPort>%d</NewExternalPort>
			</u:DeletePortMapping>
		</s:Body>
	</s:Envelope>
	`
	deletePortMappingUrl = `"urn:schemas-upnp-org:service:WANIPConnection:1#DeletePortMapping"`
)

type UPNPForwarder struct {
	SsdpPort  int
	InnerPort int
	OuterPort int

	upnpInterfaces map[string]bool
}

type UPNPInformation struct {
	Key string
	Value string
}
