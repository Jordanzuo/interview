package command

import (
	"encoding/json"
	"fmt"

	"interview.com/cloudcade/chat/client/src/model"
)

type SendMessageParameter struct {
	Message string
}

type SendMessageResponseData struct {
	PlayerName string
	Message    string
}

func SendMessage(requestID int64, paramList []string) (requestObj *model.RequestObject, callback func(interface{}), err error) {
	if len(paramList) != 2 {
		err = fmt.Errorf("Format: %s Message", Command_sendMessage)
		return
	}

	parameter := &SendMessageParameter{
		Message: paramList[1],
	}
	requestObj = model.NewRequestObject(requestID, Command_sendMessage, parameter)
	callback = sendMessageCallback
	return
}

func sendMessageCallback(data interface{}) {

}

func SendMessagePush(data interface{}) {
	var sendMessageResponseDataObj *SendMessageResponseData
	bytes, _ := json.Marshal(data)
	err := json.Unmarshal(bytes, &sendMessageResponseDataObj)
	if err != nil {
		fmt.Printf("Unmarshal send message push data failed. Error:%s\n", err)
		return
	}

	fmt.Printf("[%s] says: %s\n", sendMessageResponseDataObj.PlayerName, sendMessageResponseDataObj.Message)
}
