package main

import (
	"fmt"
	"sync"

	_ "interview.com/cloudcade/chat/server/src/clientmgr"
	"interview.com/cloudcade/chat/server/src/config"
	"interview.com/cloudcade/chat/server/src/dfa"
	_ "interview.com/cloudcade/chat/server/src/player"
	"interview.com/cloudcade/chat/server/src/room"
	"interview.com/cloudcade/chat/server/src/serversocket"
	"interview.com/cloudcade/chat/server/src/serverwebsocket"
)

var (
	wg sync.WaitGroup
)

func init() {
	wg.Add(2)
}

func main() {
	fmt.Println("Hello chat server.")

	err := config.Init("config/config.json")
	if err != nil {
		panic(fmt.Sprintf("Init config failed. Error:%s", err))
	}

	err = dfa.Init("config/profanity_words.txt")
	if err != nil {
		panic(fmt.Sprintf("Init profanity filter logic failed. Error:%s", err))
	}

	// Init room data
	configObj := config.GetConfig()
	room.Init(configObj.RoomCount)

	// Start socket server
	go serversocket.Start(&wg, configObj.SocketServerListenAddr)

	// Start web socket server
	go serverwebsocket.Start(&wg, configObj.WebSocketServerListenAddr)

	wg.Wait()
}
