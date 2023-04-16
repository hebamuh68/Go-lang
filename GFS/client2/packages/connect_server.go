package GFS

import (
	"net"
)

func Create_Port(network string, address string){
	ln, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}

	defer ln.Close()
}