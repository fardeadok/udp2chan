package udp2chan

import (
	"errors"
	"fmt"
	"log"
	"net"
)

func err(err error) string {
	if err != nil {
		log.Fatal()
		return err.Error()
	}
}

//Udpserver is struct for  udp -> chan
type Udpserver struct {
	name       string
	addr, port string
	network    string
	udpaddr    *net.UDPAddr
	conn       *net.UDPConn
	ch         chan int
}

//Start - start listening and transmit to chan
func (v Udpserver) Start() {
	defer v.conn.Close()
	for {
		buffer := make([]byte, 1024)
		n, addr, err := v.conn.ReadFromUDP(buffer)
		if err != nil {
		}
		fmt.Println("udp client:", addr)
		fmt.Print("received:", string(buffer[:n]))
	}
}

//New -  (network name,  port)
func New(name, addr, port string) (*Udpserver, error) {

	if addr == "" || port == "" {
		return nil, errors.New("wrong addr:port")
	}

	tmpAddr, err := net.ResolveUDPAddr("udp4", addr+":"+port)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tmpConn, err := net.ListenUDP("udp4", tmpAddr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tmpServ := &Udpserver{
		name:    name,
		port:    port,
		ch:      make(chan int, 1000),
		udpaddr: tmpAddr,
		network: "udp4",
		conn:    tmpConn,
	}

	return tmpServ, nil
}
