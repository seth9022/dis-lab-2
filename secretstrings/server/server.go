package main

import (
	"dis-lab-2/secretstrings/stubs"
	"flag"
	"net"
	"net/rpc"

	//	"errors"
	//	"flag"
	//	"fmt"
	//	"net"
	"math/rand"
	"time"
	//	"secretstrings/stubs"
	//	"net/rpc"
)

/** Super-Secret `reversing a string' method we can't allow clients to see. **/
func ReverseString(s string, i int) string {
	time.Sleep(time.Duration(rand.Intn(i)) * time.Second)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type SecretStringOperations struct{}

func (s *SecretStringOperations) Reverse(req stubs.Request, res *stubs.Request) (err error) {
	res.Message = ReverseString(req.Message, 10) // i sets the time
	return
}

func (s *SecretStringOperations) FastReverse(req stubs.Request, res *stubs.Request) (err error) {
	res.Message = ReverseString(req.Message, 2)
	return
}

func main() {
	pAddr := flag.String("port", "8030", "Port to listen on")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	rpc.Register(&SecretStringOperations{})
	listener, _ := net.Listen("tcp", ":"+*pAddr)
	defer listener.Close()
	rpc.Accept(listener)

}
