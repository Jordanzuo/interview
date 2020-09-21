package clientmgr

import (
	"fmt"

	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

var (
	loginFuncMap = make(map[string]func(*model.RequestObject, IClient) *model.ResponseObject, 2)
	funcMap      = make(map[string]func(*model.RequestObject, IClient, *playerModel.Player) *model.ResponseObject, 16)
)

// RegisterLoginFunction ... Other modules register their login related functions to loginFuncMap
// command: Command of the function
// funcItem: The target function
func RegisterLoginFunction(command string, funcItem func(*model.RequestObject, IClient) *model.ResponseObject) {
	if _, exists := loginFuncMap[command]; exists {
		panic(fmt.Sprintf("%s:Function:%s has been registered.", serverModuleName, command))
	}

	loginFuncMap[command] = funcItem
}

func getLoginFuncItem(command string) (funcItem func(*model.RequestObject, IClient) *model.ResponseObject, exists bool) {
	funcItem, exists = loginFuncMap[command]
	return
}

// RegisterFunction ... Other modules register their functions to funcMap
// command: Command of the function
// funcItem: The target function
func RegisterFunction(command string, funcItem func(*model.RequestObject, IClient, *playerModel.Player) *model.ResponseObject) {
	if _, exists := funcMap[command]; exists {
		panic(fmt.Sprintf("%s:Function:%s has been registered.", serverModuleName, command))
	}

	funcMap[command] = funcItem
}

func getFuncItem(command string) (funcItem func(*model.RequestObject, IClient, *playerModel.Player) *model.ResponseObject, exists bool) {
	funcItem, exists = funcMap[command]
	return
}
