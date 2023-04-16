package GFS

import (
	"fmt"
	"net"
	"os"
)


func Connect_G_slave(slaveAddr string) {

		//=========== 1) Connect with the slave
		slave, err := net.Dial("tcp", "127.0.0.1:8088")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer slave.Close()


		////=========== 2) Send to the slave the client port which needs data
		_, err = slave.Write([]byte(slaveAddr))
		if err != nil {
			fmt.Println(err)
			return
		}

	}

