package GFS

import (
	"fmt"
	"net"
	"os"
)


func Connect_slave(clientAddr string) {

	for {

		//Connect with the slave
		slave, err := net.Dial("tcp", "127.0.0.1:8090")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer slave.Close()


		//Send to the slave the port which needs data
		_, err = slave.Write([]byte("hi i'm slave"))
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
