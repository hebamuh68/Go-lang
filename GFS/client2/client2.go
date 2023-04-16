package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Starting client 2...")


	// Create port
	client2, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer client2.Close()


	// Connect with server
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buffer[:n]))

}
