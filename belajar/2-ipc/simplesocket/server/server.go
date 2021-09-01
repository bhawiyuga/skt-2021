package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":9999")
	errorHandling(err)

	listener, err := net.ListenTCP("tcp", addr)
	errorHandling(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		errorHandling(err)
		/*
			data := make([]byte, 256)

			n, err := conn.Read(data)
			if n > 0 {
				fmt.Println(n, string(data[:n]))
			}*/
		data := make([]byte, 256)
		_, err = conn.Read(data)
		errorHandling(err)
		fmt.Println(string(data))

		dataInStr := string(data)
		dataInStr = "OK " + dataInStr
		conn.Write([]byte(dataInStr))
	}
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
}
