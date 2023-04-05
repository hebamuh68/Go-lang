package main

import (

	"fmt"
	"net"
	"gfs.com/server/packages"
	"strings"
)

func main() {

	fmt.Println("Starting server...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Get what data user want
		msg := strings.TrimSpace(string(buf[:n]))
		fmt.Println("\nUser need data of student:", msg)


		//Send the data to the user
		if strings.Contains(msg, "Names") {
			
			response := GFS.Read_Data("Student_Grades.txt")
			conn.Write([]byte(response))
			fmt.Println("Sent response to client:", response)
		}		
	}
}



