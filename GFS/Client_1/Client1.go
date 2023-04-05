package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting client...")

	// Replace the IP address and port with the address of your server
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	for {

		//Send which data user need
		fmt.Print("Enter Student data you need: ")
		var msg string
		fmt.Scanln(&msg)

		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		response := string(buf[:n])
		fmt.Println("Received response:", response)
	}
}
