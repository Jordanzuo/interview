package clientmgr

import (
	"encoding/json"
	"fmt"

	"interview.com/cloudcade/chat/server/src/model"
)

const conLoginCommand = "/Login"

// HandleRequest ... Handle request from client
// clientObj：client object
// request： request content
// Return values:
// Non
func HandleRequest(clientObj IClient, request []byte) {
	var requestObj *model.RequestObject
	var responseObj = model.NewResponseObject()
	var command string

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[%s]HandleRequest encounters unhandled error:%v", serverModuleName, r)
			responseObj.SetResponseStatus(model.RuntimeException)
		}

		clientObj.SendMessage(responseObj)

		if command == conLoginCommand {
			if responseObj.ResponseStatus != model.Success {
				clientObj.SendMessage(model.NewDisconnectResponseObject())
			}
		}
	}()

	// Unmarshal request data
	requestObj = new(model.RequestObject)
	if err := json.Unmarshal(request, requestObj); err != nil {
		fmt.Printf("Unmarshal request failed. Error:%s\n", err)
		return
	}
	command = requestObj.Command

	// Check if the request id is incrementing
	responseObj.RequestID = requestObj.RequestID
	if requestObj.RequestID <= clientObj.GetClientRequestID() {
		responseObj.SetResponseStatus(model.ClientRequestExpired)
		return
	}
	clientObj.SetClientRequestID(requestObj.RequestID)

	// Choose the right function to execute
	if funcItem, exists := getLoginFuncItem(command); exists {
		responseObj = funcItem(requestObj, clientObj)
		responseObj.RequestID = requestObj.RequestID
	} else if funcItem, exists := getFuncItem(command); exists {
		playerObj := clientObj.GetPlayer()
		if playerObj == nil {
			responseObj.SetResponseStatus(model.NoLogin)
			return
		}

		responseObj = funcItem(requestObj, clientObj, playerObj)
		responseObj.RequestID = requestObj.RequestID
	} else {
		responseObj.SetResponseStatus(model.NoTargetMethod)
	}
}
