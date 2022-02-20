package kademlia

import (
	"encoding/hex"
	"fmt"
	"github.com/kk222mo/godist/config"
	"math/rand"
	"time"
)

func (k Key) GetBit(position int) int {
	rand.Seed(time.Now().UnixNano())
	bytenum := position / 8
	bitnum := position % 8
	if k.Data[bytenum]&(1<<(8-1-bitnum)) != 0 {
		return 1
	}
	return 0
}

func (k Key) ToHexString() string {
	res := ""
	for i := 0; i < config.KEY_SIZE; i++ {
		res += fmt.Sprintf("%02x", k.Data[i])
	}
	return res
}
func (k Key) ToBinaryString() string {
	res := ""
	for i := 0; i < config.KEY_SIZE; i++ {
		res += fmt.Sprintf("%08b", k.Data[i])
	}
	return res
}

func KeyFromHexString(s string) (Key, error) {
	data, err := hex.DecodeString(s)
	if err != nil {
		return Key{}, err
	}
	var keyData [config.KEY_SIZE]byte
	copy(keyData[:], data)
	return Key{Data: keyData}, nil
}

func GenerateRandomKey() Key {
	k := Key{}
	rand.Read(k.Data[:])
	return k
}

func Distance(k1, k2 Key) Key {
	var res Key
	for i := 0; i < len(k1.Data); i++ {
		res.Data[i] = k1.Data[i] ^ k2.Data[i]
	}
	return res
}
