package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve alamat IP dan Port dari server
	serverAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9000")
	errorHandling(err)

	// Kirim permintaan koneksi TCP
	conn, err := net.DialTCP("tcp", nil, serverAddr)
	errorHandling(err)

	// Kirim string ke server
	data := "Selamat pagi"
	conn.Write([]byte(data))

	// Baca response dari server
	buff := make([]byte, 256)
	_, err = conn.Read(buff)
	fmt.Println("Menerima respon dari server : ", string(buff))

	// Tutup koneksi
	conn.Close()
}

func errorHandling(err error) {
	if err != nil {
		fmt.Println("Terjadi error", err.Error())
	}
}
