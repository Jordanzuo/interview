package rpc

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"

	"interview.com/cloudcade/chat/client/src/command"
	"interview.com/cloudcade/chat/client/src/model"
)

var (
	clientObj *client
)

func handleClient() {
	for {
		content, ok := clientObj.getValidMessage()
		if !ok {
			break
		}

		var responseObj *model.ResponseObject
		err := json.Unmarshal(content, &responseObj)
		if err != nil {
			fmt.Sprintf("Unmarshal %s failed. Error:%s\n", string(content), err)
			continue
		}

		if responseObj.ResponseStatus != "" {
			fmt.Printf("ResponseStatus: %s\n", responseObj.ResponseStatus)
			continue
		}

		// Handle push data
		if responseObj.PushKey != "" {
			switch responseObj.PushKey {
			case model.SendMessage:
				command.SendMessagePush(responseObj.Data)
			case model.LoginOnAnotherDevice:
				fmt.Println("Your account is logining on another device.")
			default:
				fmt.Println("Undefined push key:%s\n", responseObj.PushKey)
			}
		} else {
			callback, exists := getCallback(responseObj.RequestID)
			if !exists {
				fmt.Println("No callback found for RequestID:%d\n", responseObj.RequestID)
			} else {
				callback(responseObj.Data)
			}
		}
	}
}

func ConnectServer(serverAddr string, ch chan int) {
	conn, err := net.DialTimeout("tcp", serverAddr, 2*time.Second)
	if err != nil {
		fmt.Printf("Dial Error: %s\n", err)
		panic("Connect to remote server failed.")
		return
	}

	fmt.Printf("Connect to the server. (local address: %s)\n", conn.LocalAddr())
	clientObj = newClient(conn)
	ch <- 1

	defer func() {
		conn.Close()
		clientObj = nil
	}()

	for {
		readBytes := make([]byte, 1024)
		n, err := conn.Read(readBytes)
		if err != nil {
			fmt.Printf("Read message failed. Error：%s\n", err)
			os.Exit(1)
		}

		clientObj.appendContent(readBytes[:n])
		handleClient()
	}
}
