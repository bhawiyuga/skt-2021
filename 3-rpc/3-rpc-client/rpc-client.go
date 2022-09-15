package main

import (
	"fmt"
	"net/rpc"
)

type MyArgs struct {
	A, B int
}

func main() {

	// Inisiasi koneksi HTTP dari client ke server
	client, err := rpc.DialHTTP("tcp", "18.140.224.119:1234")
	handleError(err)
	// Argument yang akan dikirimkan dari client ke server
	clientargs := &MyArgs{1000, 500}
	// Pointer untuk menampung hasil eksekusi dari server
	var result int
	err = client.Call("Arith.Substract", clientargs, &result)
	handleError(err)
	fmt.Println("Hasil eksekusi RPC server : ", result)

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Terdapat error : ", err.Error())
	}
}
