package main

import (
	"fmt"
	"os"

	"interview.com/cloudcade/chat/client/src/rpc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./client ip:host")
		return
	}

	serverAddr := os.Args[1]

	startCh := make(chan int)
	go rpc.ConnectServer(serverAddr, startCh)
	<-startCh

	ch := make(chan int)
	go rpc.Interact(ch)

	<-ch
}
