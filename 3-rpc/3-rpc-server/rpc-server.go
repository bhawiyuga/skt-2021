package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
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

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Terdapat error : ", err.Error())
	}
}
