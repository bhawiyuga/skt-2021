package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Bagian untuk resolve IP address dan port dari server
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9000")
	handleError(err)

	// Listen ke sebuah port dengan protokol TCP sebagai layer transportnya
	listener, err := net.ListenTCP("tcp", addr)
	handleError(err)

	for {
		// Terima permintaan koneksi dari client
		conn, err := listener.Accept()
		handleError(err)

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// Baca data dari client
	// - Membuat string buffer berukuran tertentu misalnya 256 byte
	buffer := make([]byte, 256)
	_, err := conn.Read(buffer)
	handleError(err)

	// Kita tambahkan string OK di depan data dari client tersebut
	dataFromClient := string(buffer)
	fmt.Println("Menerima data dari client : ", dataFromClient)
	dataFromClient = "OK " + dataFromClient

	// Kirimkan kembali ke client
	conn.Write([]byte(dataFromClient))
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error message : ", err.Error())
		os.Exit(1)
	}
}
