package clientmgr

import (
	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

func pushToPlayer(playerObj *playerModel.Player, responseObj *model.ResponseObject) {
	if playerObj.ClientID == 0 {
		return
	}

	clientObj, exists := GetClient(playerObj.ClientID)
	if !exists {
		return
	}

	clientObj.SendMessage(responseObj)
}

func PushMessageToPlayer(playerObj *playerModel.Player, pushKey string, message interface{}) {
	sendResponseObj := model.NewResponseObject()
	sendResponseObj.SetPushKey(pushKey)
	sendResponseObj.SetData(message)

	pushToPlayer(playerObj, sendResponseObj)
}

func PushMessageToPlayerList(playerList []*playerModel.Player, pushKey string, message interface{}) {
	sendResponseObj := model.NewResponseObject()
	sendResponseObj.SetPushKey(pushKey)
	sendResponseObj.SetData(message)

	for _, item := range playerList {
		pushToPlayer(item, sendResponseObj)
	}
}
