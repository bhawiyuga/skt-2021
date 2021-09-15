package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type MyArgs struct {
	A, B int
}

type Arith struct {
}

func (t *Arith) Add(args *MyArgs, result *int) error {
	*result = args.A + args.B
	return nil
}

func main() {
	// Menangkap parameter client/server dari user
	args := os.Args[1]
	if args == "server" {
		// Inisiasi struct arith
		arith := &Arith{}
		// Registrasikan struct dan method ke RPC
		rpc.Register(arith)
		// Deklarasikan bahwa kita menggunakan protokol HTTP sebagai mekanisme pengiriman pesan
		rpc.HandleHTTP()
		// Deklarasikan listerner HTTP dengan layer transport TCP dan Port 1234
		listener, err := net.Listen("tcp", ":1234")
		handleError(err)
		// Jalankan server HTTP
		http.Serve(listener, nil)
	} else if args == "client" {
		// Inisiasi koneksi HTTP dari client ke server
		client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
		handleError(err)
		// Argument yang akan dikirimkan dari client ke server
		clientargs := &MyArgs{1000, 2000}
		// Pointer untuk menampung hasil eksekusi dari server
		var result int
		err = client.Call("Arith.Add", clientargs, &result)
		handleError(err)
		fmt.Println("Hasil eksekusi RPC server : ", result)
	}

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Terdapat error : ", err.Error())
	}
}
