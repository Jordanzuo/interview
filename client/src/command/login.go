package command

import (
	"encoding/json"
	"fmt"

	"interview.com/cloudcade/chat/client/src/model"
)

type LoginParameter struct {
	Name string
}

type LoginResponseData struct {
	MessageHistoryList []*Message
}

type Message struct {
	PlayerName string
	Message    string
}

func Login(requestID int64, paramList []string) (requestObj *model.RequestObject, callback func(interface{}), err error) {
	if len(paramList) != 2 {
		err = fmt.Errorf("Format: %s PlayerName", Command_login)
		return
	}

	parameter := &LoginParameter{
		Name: paramList[1],
	}
	requestObj = model.NewRequestObject(requestID, Command_login, parameter)
	callback = loginCallback
	return
}

func loginCallback(data interface{}) {
	fmt.Println("Login successfully!")
	var responseData *LoginResponseData

	bytes, _ := json.Marshal(data)
	err := json.Unmarshal(bytes, &responseData)
	if err != nil {
		fmt.Printf("Unmarshal login result failed. Error:%s\n", err)
		return
	}

	for _, item := range responseData.MessageHistoryList {
		fmt.Printf("[%s] says: %s\n", item.PlayerName, item.Message)
	}
}
