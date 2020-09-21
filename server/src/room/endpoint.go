package room

import (
	"interview.com/cloudcade/chat/server/src/clientmgr"
	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
	"interview.com/cloudcade/chat/server/src/popular"
)

func init() {
	clientmgr.RegisterFunction("/sendMessage", sendMessage)
}

type SendMessageParameter struct {
	Message string
	roomObj *Room
}

func (this *SendMessageParameter) verify(playerObj *playerModel.Player) model.ResponseStatus {
	if this.Message == "" {
		return model.MessageIsEmpty
	}

	roomID := playerObj.RoomID
	if roomID == 0 {
		return model.PlayerNotInRoom
	}

	var exists bool
	this.roomObj, exists = getRoom(roomID)
	if !exists {
		return model.PlayerNotInValidRoom
	}

	return model.Success
}

func sendMessage(requestObj *model.RequestObject, clientObj clientmgr.IClient, playerObj *playerModel.Player) *model.ResponseObject {
	var responseObj = model.NewResponseObject()
	var paramObj = new(SendMessageParameter)
	var rs model.ResponseStatus

	rs = requestObj.ParseParameter(&paramObj)
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	rs = paramObj.verify(playerObj)
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	paramObj.roomObj.AppendMessage(playerObj, paramObj.Message)

	// Send message to popular
	popular.AddMessage(paramObj.Message)

	// Push message to all the players in the same room
	playerList := paramObj.roomObj.GetAllPlayers()
	clientmgr.PushMessageToPlayerList(playerList, model.SendMessage, newRoomMessage(playerObj.Name, paramObj.Message))

	return responseObj
}
