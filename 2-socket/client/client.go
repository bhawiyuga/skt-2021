package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Resolve address
	serverAddress, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9000")
	handleError(err)
	// Kirim permintaan koneksi ke server
	conn, err := net.DialTCP("tcp", nil, serverAddress)
	handleError(err)

	// Kirim data ke server
	data := "Selamat pagi"
	_, err = conn.Write([]byte(data))
	handleError(err)

	// Baca respon data dari server
	responseData := make([]byte, 256)
	_, err = conn.Read(responseData)
	handleError(err)

	// Cetak reponse dari server
	responseDataInStr := string(responseData)
	fmt.Println(responseDataInStr)

	// Tutup koneksinya
	conn.Close()

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error message : ", err.Error())
		os.Exit(1)
	}
}
