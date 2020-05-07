package udp2chan

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

var (
	globalE = 100
)

//init - is auto start
func init() {
	globalE = 0
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

//UDPServer is struct for  udp -> chan
type UDPServer struct {
	name       string
	addr, port string
	network    string
	udpaddr    *net.UDPAddr
	conn       *net.UDPConn
	ch         chan string
	sync.RWMutex
}

//UDPServers  - []* UdpServer
type UDPServers struct {
	Map map[string]*UDPServer // string is  name
}

//Start - start goroutine listening and transmit to chan
func (v UDPServer) Start() {
	v.Lock()
	defer v.conn.Close()
	defer v.Unlock()
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
func New(name, addr, port string) (*UDPServer, error) {

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

	tmpServ := &UDPServer{
		name:    name,
		port:    port,
		ch:      make(chan string, 1000),
		udpaddr: tmpAddr,
		network: "udp4",
		conn:    tmpConn,
	}

	return tmpServ, nil
}
