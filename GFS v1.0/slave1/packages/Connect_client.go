package GFS

import (
	"fmt"
	"net"
	"os"
)


func Connect_client(client_addr string) {

	//=========== 1) connect client
    client, err := net.Dial("tcp","127.0.0.1:8090")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer client.Close()


	//=========== 1) send the data to client
	
	//Send the data to the user
	response := Read_Data("Student_Grades.txt")
	client.Write([]byte(response))
	fmt.Println("Sending response...")
	
}