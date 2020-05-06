package main

import (
	"fmt"

	"github.com/fardeadok/udp2chan"
)

func main() {
	// v001 := udp2chan.Udpserver{}
	// "udp4"
	v002, err := udp2chan.New("serv001","localhost", "9000")
	if err != nil {
	return
	}

	fmt.Println(v002)

	return
}
