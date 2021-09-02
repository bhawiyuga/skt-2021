package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve alamat IP dan Port dari server
	// - Server akan diikat pada alamat IP 0.0.0.0 (tidak ada batasan) dan port 9000
	address, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9000")
	errorHandling(err)

	// Listen dengan protokol TCP sebagai layer transportnya
	listener, err := net.ListenTCP("tcp", address)
	errorHandling(err)

	for {
		// Menerima permintaan koneksi dari client
		conn, err := listener.Accept()
		errorHandling(err)

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Baca data dari client
	// - Deklarasikan buffer untuk menampung data
	buff := make([]byte, 256)
	_, err := conn.Read(buff)
	errorHandling(err)
	fmt.Println("Menerima data dari client : ", string(buff))
	// - Casting dari slice of byte ke string
	data := string(buff)
	data = "OK " + data

	// Kirim balik ke client
	conn.Write([]byte(data))
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println("Terjadi error", err.Error())
	}
}
