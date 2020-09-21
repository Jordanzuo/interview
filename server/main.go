package main

import (
	"fmt"

	"interview.com/cloudcade/chat/server/src/config"
)

func main() {
	fmt.Println("Hello chat server.")

	config.Init("config/config.json")
}
