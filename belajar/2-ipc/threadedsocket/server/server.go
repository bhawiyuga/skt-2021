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

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	data := make([]byte, 256)
	_, err := conn.Read(data)
	errorHandling(err)
	fmt.Println(string(data))

	dataInStr := string(data)
	dataInStr = "OK " + dataInStr
	conn.Write([]byte(dataInStr))

	conn.Close()
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
}
