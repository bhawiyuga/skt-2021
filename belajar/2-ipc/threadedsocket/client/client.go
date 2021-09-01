package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	errorHandling(err)

	conn, err := net.DialTCP("tcp", nil, addr)
	errorHandling(err)

	data := "Selamat pagi"
	conn.Write([]byte(data))
	fmt.Println("data sent")

	dataReturn := make([]byte, 256)
	_, err = conn.Read(dataReturn)
	errorHandling(err)
	fmt.Println(string(dataReturn))

	conn.Close()
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}
