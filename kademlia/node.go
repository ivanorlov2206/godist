package kademlia

/*import (
	"fmt"
)*/

func PushNode(head *Vertex, n Node) {
	v := FindNearestBucket(head, n.Id)

	(*v).addNode(n)
}
