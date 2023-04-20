package GFS

import (
	"fmt"
	"net"
	"os"
)


func Connect_master(data string) {

    //=========== 1) connect master
    master, err := net.Dial("tcp", "127.0.0.1:8080")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer master.Close()


    //=========== 2) send name of data needed
    master.Write([]byte(data))
    fmt.Println("Sent request to get:", data)
}
