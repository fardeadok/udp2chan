package udpserv

import (
	"fmt"
	"net"
)

var (
//TestVar001 int = 100
)

//
//
//UDP2chan is struct for  udp -> chan
type UDP2chan struct {
	name    string
	port    string
	network string
	udpaddr *net.UDPAddr
	conn    *net.UDPConn
	ch      chan int
}

//
//
//
func (v UDP2chan) String() string {
	return v.name
}

//
//
//
func (v UDP2chan) GetName() string {
	return v.name
}

//
//
//NewUDP2chan -  name  port
func NewUDP2chan(name, port string) (*UDP2chan, error) {
	tmp := &UDP2chan{
		name: name,
		port: port,
		ch:   make(chan int, 1000),
	}

	tmpU, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tmp.udpaddr = tmpU

	tmpConn, err := net.ListenUDP("udp4", tmpU)
	tmp.conn = tmpConn

	return tmp, nil
}
