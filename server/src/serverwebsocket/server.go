package serverwebsocket

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"interview.com/cloudcade/chat/server/src/clientmgr"
)

var (
	isServerClosing = false
	upgrader        = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleConn(w http.ResponseWriter, r *http.Request) {
	if isServerClosing {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	clientObj := newClient(conn)
	clientObj.start()
	clientmgr.RegisterClient(clientObj)
}

// Start ... Start web socket server
func Start(wg *sync.WaitGroup, address string) {
	defer wg.Done()

	msg := fmt.Sprintf("serverwebsocket begins to listen on:%s...", address)
	fmt.Println(msg)

	http.HandleFunc("/", handleConn)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(fmt.Sprintf("serverwebsocket.ListenAndServe, err:%v", err))
	}
}
