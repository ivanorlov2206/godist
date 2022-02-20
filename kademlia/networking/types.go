package networking

import (
	"github.com/kk222mo/godist/kademlia"
)

type NodeMessage struct {
	RequestType string
	FromID      kademlia.Key
	Headers     map[string]string
	Body        string
}
