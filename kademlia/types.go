package kademlia

import (
	"github.com/kk222mo/godist/config"
)

type Key struct {
	Data [config.KEY_SIZE]byte
}

type Node struct {
	Id      Key
	Address string
}

type KBucket struct {
	Nodes []Node
}

type Vertex struct { // Node struct is binary tree
	PrefixLen int
	Bucket    KBucket
	Parent    *Vertex
	Left      *Vertex // 1
	Right     *Vertex // 0
}
