package rpc

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interview.com/cloudcade/chat/client/src/command"
	"interview.com/cloudcade/chat/client/src/model"
)

func Interact(ch chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Interactive encounters unhandled error. r:%v\n", r)
		}

		ch <- 1
	}()

	fmt.Println("Please type into messages, q to quit.")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		input := scanner.Text()

		if input == "" {
			fmt.Println("Please type into messages, q to quit.")
			continue
		}

		if input == "q" {
			clientObj.conn.Close()
			break
		}

		requestObj, callback, err := constructRequestObject(input)
		if err != nil {
			fmt.Println(err)
			return
		}
		request(requestObj, callback)
	}
}

func constructRequestObject(message string) (requestObj *model.RequestObject, callback func(interface{}), err error) {
	msgList := strings.SplitN(message, " ", 2)
	if len(msgList) < 1 {
		err = fmt.Errorf("Format: command parameter")
		return
	}

	switch msgList[0] {
	case command.Command_login:
		requestObj, callback, err = command.Login(generateNewRequestID(), msgList)
	case command.Command_popular:
		requestObj, callback, err = command.Popular(generateNewRequestID(), msgList)
	case command.Command_sendMessage:
		requestObj, callback, err = command.SendMessage(generateNewRequestID(), msgList)
	case command.Command_stats:
		requestObj, callback, err = command.Stats(generateNewRequestID(), msgList)
	default:
		err = fmt.Errorf("Unknown command:%s", msgList[0])
		return
	}

	return
}
