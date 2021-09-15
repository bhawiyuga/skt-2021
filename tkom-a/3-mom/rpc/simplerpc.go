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
	args := os.Args[1]

	if args == "server" {
		arith := &Arith{}
		rpc.Register(arith)
		rpc.HandleHTTP()
		listener, err := net.Listen("tcp", ":1234")
		handleError(err)
		http.Serve(listener, nil)
	} else {
		client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
		handleError(err)
		var result int
		err = client.Call("Arith.Add", &MyArgs{20, 20}, &result)
		handleError(err)
		fmt.Println("Result : ", result)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
}
