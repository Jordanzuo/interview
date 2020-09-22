package room

import (
	"fmt"

	"interview.com/cloudcade/chat/server/src/dfa"

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

	this.Message = dfa.HandleWord(this.Message, '*')

	return model.Success
}

type SendMessageResponseData struct {
	PlayerName string
	Message    string
}

func newSendMessageResponseData(playerName, message string) *SendMessageResponseData {
	return &SendMessageResponseData{
		PlayerName: playerName,
		Message:    message,
	}
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
	fmt.Printf("There are %d players in room:%d\n", len(playerList), paramObj.roomObj.ID)
	sendMessageResponseDataObj := newSendMessageResponseData(playerObj.Name, paramObj.Message)
	clientmgr.PushMessageToPlayerList(playerList, model.SendMessage, sendMessageResponseDataObj)

	return responseObj
}
