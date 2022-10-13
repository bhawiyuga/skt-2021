package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

var result int = 1

type MyArgs struct {
	A int
}

type Arith struct {
}

func (t *Arith) Ping(args *MyArgs, result *int) error {
	*result = args.A + 1
	return nil
}

func ping(peerAddress string) {
	fmt.Println("Try to connect to : " + peerAddress)
	// Inisiasi koneksi HTTP dari client ke server
	client, err := rpc.DialHTTP("tcp", peerAddress)
	handleConnError(err, peerAddress)

	if err == nil {
		for true {
			// Argument yang akan dikirimkan dari client ke server
			clientargs := &MyArgs{result}
			fmt.Println("Ping : ", result)
			// Pointer untuk menampung hasil eksekusi dari server
			var localResult int
			err = client.Call("Arith.Ping", clientargs, &localResult)
			result = localResult
			handleConnError(err, peerAddress)
			fmt.Println("Pong : ", result)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	portNumber := os.Args[1]
	peersAddress := os.Args[2]

	//time.Sleep(10 * time.Second)

	go ping(peersAddress)
	// Inisiasi struct arith
	arith := &Arith{}
	// Registrasikan struct dan method ke RPC
	rpc.Register(arith)
	// Deklarasikan bahwa kita menggunakan protokol HTTP sebagai mekanisme pengiriman pesan
	rpc.HandleHTTP()
	// Deklarasikan listerner HTTP dengan layer transport TCP dan Port 1234
	listener, err := net.Listen("tcp", ":"+string(portNumber))
	handleError(err)
	// Jalankan server HTTP
	http.Serve(listener, nil)

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Terdapat error : ", err.Error())
	}
}

func handleConnError(err error, peersAdrress string) {
	if err != nil {
		fmt.Println("Terdapat error : ", err.Error())
		time.Sleep(5 * time.Second)
		ping(peersAdrress)
	}

}
