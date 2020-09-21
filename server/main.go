package main

import (
	"fmt"

	"interview.com/cloudcade/chat/server/src/config"
	"interview.com/cloudcade/chat/server/src/dfa"
)

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
}
