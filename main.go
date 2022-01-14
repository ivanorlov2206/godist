package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/kk222mo/godist/portforwarder"
	"github.com/kk222mo/godist/kademlia"
)

func dfs(v *kademlia.Vertex, level int, path []string) {
	if v == nil {
		return
	}

	fmt.Printf("On level %d [%s] with prefix len %d we have %d nodes\n", level, strings.Join(path, ", "), (*v).PrefixLen, len((*v).Bucket.Nodes))
	dfs((*v).Left, level+1, append(path, "left"))
	dfs((*v).Right, level+1, append(path, "right"))
}

func main() {
	fmt.Println("Starting port forwarding")
	forwarder := portforwarder.NewUPNPForwarder(22222, 33334, 33334)
	readyStream := make(chan string)
	portforwarder.StartForwarding("tcp", forwarder, readyStream)
	<-readyStream
	for {

	}
	v := kademlia.Vertex{PrefixLen: 0, Bucket: kademlia.KBucket{}}
	fmt.Println(v)
	f, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i < 20; i++ {
		k1 := kademlia.GenerateRandomKey()
		fmt.Fprintln(f, k1.ToBinaryString())
		n1 := kademlia.Node{Id: k1}
		kademlia.PushNode(&v, n1)
	}
	dfs(&v, 0, []string{""})
}
