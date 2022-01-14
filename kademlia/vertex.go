package kademlia

import (
	"github.com/kk222mo/godist/config"
)

func (vertex *Vertex) Split() {
	left, right := Vertex{Bucket: KBucket{}, PrefixLen: vertex.PrefixLen + 1,
		Left: nil, Right: nil, Parent: vertex}, Vertex{Bucket: KBucket{}, PrefixLen: vertex.PrefixLen + 1,
		Left: nil, Right: nil, Parent: vertex}
	for i := 0; i < len(vertex.Bucket.Nodes); i++ {
		node := vertex.Bucket.Nodes[i]
		if node.Id.GetBit(vertex.PrefixLen) == 1 {
			left.Bucket.Nodes = append(left.Bucket.Nodes, node)
		} else {
			right.Bucket.Nodes = append(right.Bucket.Nodes, node)
		}
	}
	vertex.Bucket.Nodes = nil
	vertex.Left = &left
	vertex.Right = &right
}

func (vertex *Vertex) addNode(n Node) {
	vertex.Bucket.Nodes = append(vertex.Bucket.Nodes, n)
	if len(vertex.Bucket.Nodes) > config.K {
		vertex.Split()
	}
}

func FindNearestBucket(head *Vertex, k Key) *Vertex {
	for i := 0; i < config.KEY_SIZE; i++ {
		if k.GetBit(i) == 1 && (*head).Left != nil {
			head = (*head).Left
		} else if k.GetBit(i) == 0 && (*head).Right != nil {
			head = (*head).Right
		} else {
			break
		}
	}
	return head
}
