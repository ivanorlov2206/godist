package kademlia

func Distance(n1, n2 Node) [256]byte {
	var res [256]byte
	for i := 0; i < len(n1.Id); i++ {
		res[i] = n1.Id[i] ^ n2.Id[i]
	}
	return res
}
