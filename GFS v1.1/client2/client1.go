package main

import (
	"fmt"
	"net"
	"gfs.com/client/packages"
)

func main() {

	//======================================= Create client
	fmt.Println("Starting client 1...")
	client1, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer client1.Close()

	//======================================= Enter data needed
	for {
		// Prompt user for input
		fmt.Printf("Enter the data needed (type 'exit' to quit): ")
		var data string
		fmt.Scanln(&data)

		// Check if user entered exit code
		if data == "exit" {
			fmt.Println("Exiting...")
			break
		}

		//======================================= Send request to master
		GFS.Connect_master(data)

		//======================================= Recieve data from slave
		conn, err := client1.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Read the response from the Slave
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Send to the client
		Recieve_data := string(buf[:n])
		fmt.Println("Recieved data:\n", Recieve_data)

	}

}
