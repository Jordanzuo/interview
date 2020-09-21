package serversocket

import (
	"fmt"
	"net"
	"sync"

	"interview.com/cloudcade/chat/server/src/clientmgr"
)

// Start ... Start socket server
func Start(wg *sync.WaitGroup, address string) {
	defer func() {
		wg.Done()
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	msg := fmt.Sprintf("serversocket begins to listen on:%s...", address)
	fmt.Println(msg)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(fmt.Sprintf("serversocket listen Error: %s", err))
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		clientObj := newClient(conn)
		clientObj.start()
		clientmgr.RegisterClient(clientObj)
	}
}
