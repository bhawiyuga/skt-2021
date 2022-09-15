package main

import (
	"fmt"
	"time"
)

var myTime int64

type Request struct {
	T1 int64
}

type Response struct {
	T2, T3 int64
}

type ClockSyncService struct {
}

func (t *ClockSyncService) sync(req *Request, res *Response) error {
	res.T2 = myTime
	res.T3 = myTime
	return nil
}

func clockTick() {
	for {
		myTime = myTime + 1
		fmt.Println(myTime)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	myTime = time.Now().Unix()
	fmt.Println(myTime)
	go clockTick()

	<-make(chan int)
}
