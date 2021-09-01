package main

import (
	"fmt"
	"time"
)

func run(id int) {
	fmt.Println("ID : ", id)
}

func main() {
	for i := 1; i <= 10; i++ {
		go run(i)
	}
	time.Sleep(100 * time.Millisecond)
}
