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

// PushMessageToPlayer ... Push message to specified player object
func PushMessageToPlayer(playerObj *playerModel.Player, pushKey string, message interface{}) {
	sendResponseObj := model.NewResponseObject()
	sendResponseObj.SetPushKey(pushKey)
	sendResponseObj.SetData(message)

	pushToPlayer(playerObj, sendResponseObj)
}

// PushMessageToPlayerList ... Push message to a group of player object
func PushMessageToPlayerList(playerList []*playerModel.Player, pushKey string, message interface{}) {
	sendResponseObj := model.NewResponseObject()
	sendResponseObj.SetPushKey(pushKey)
	sendResponseObj.SetData(message)

	for _, item := range playerList {
		pushToPlayer(item, sendResponseObj)
	}
}
