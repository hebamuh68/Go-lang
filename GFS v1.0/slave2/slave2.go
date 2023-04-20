package main

import (
	"fmt"
	"net"
	"gfs.com/client/packages"
	

)

func main() {


	//======================================= Create slave
	fmt.Println("Starting slave...")
	slave, err := net.Listen("tcp", "127.0.0.1:8099")
	if err != nil {
		panic(err)
	}
	defer slave.Close()


	//======================================= Recieve requests from master
	for {

		conn, err := slave.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Read the response from the Master
		buf := make([]byte, 1024)
		client_addr, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		client_address := string(buf[:client_addr])
		fmt.Println("Sending Grades data to client on port: ", client_address)

	//======================================= Send the data to the client
		GFS.Connect_client(client_address)


	}
}
