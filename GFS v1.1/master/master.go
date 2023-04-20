package main

import (

	"fmt"
	"net"
	"gfs.com/server/packages"
)

func main() {


	//======================================= Create master
	fmt.Println("Starting master...")
	master, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer master.Close()


	//======================================= Recieve requests from client
	for {
		conn, err := master.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		clientAddr := conn.RemoteAddr().String()
		fmt.Println("Accepted connection from client port: ", clientAddr)


	//======================================= Check which slave should connect to
	buf := make([]byte, 1024)
	data, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	receivedData := string(buf[:data])
	fmt.Println("Data requested: ", receivedData)

	//======================================= Connect the slaves based on requested data
	if (receivedData == "Grades") {
		go GFS.Connect_G_slave(clientAddr)

	} else if (receivedData == "Names") {
		go GFS.Connect_N_slave(clientAddr)
	}
		
	}
	
}