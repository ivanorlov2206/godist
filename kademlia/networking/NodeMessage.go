package networking

import (
	"github.com/kk222mo/godist/kademlia"
	"strings"
)

func ParseNodeMessage(messageEncoded []byte) (NodeMessage, error) {
	messageDecoded := string(messageEncoded)
	result := NodeMessage{}
	headersAndBody := strings.Split(messageDecoded, "\r\n\r\n")
	headerLines := strings.Split(headersAndBody[0], "\r\n")
	typeAndSrc := strings.Split(headerLines[0], " ")
	result.RequestType = typeAndSrc[0]
	fromID, err := kademlia.KeyFromHexString(typeAndSrc[1])
	if err != nil {
		return NodeMessage{}, err
	}
	result.FromID = fromID
	headers := map[string]string{}
	for _, val := range headerLines[1:] {
		keyAndVal := strings.Split(val, ": ")
		headers[keyAndVal[0]] = keyAndVal[1]
	}
	result.Headers = headers
	result.Body = headersAndBody[1]
	return result, nil
}
