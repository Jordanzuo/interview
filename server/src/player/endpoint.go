package player

import (
	"interview.com/cloudcade/chat/server/src/clientmgr"
	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
	"interview.com/cloudcade/chat/server/src/room"
)

func init() {
	clientmgr.RegisterLoginFunction("/login", login)
	clientmgr.RegisterLoginFunction("/stats", stats)
}

type LoginParameter struct {
	Name string
}

func (this *LoginParameter) verify() model.ResponseStatus {
	if this.Name == "" {
		return model.ParamInValid
	}

	return model.Success
}

func login(requestObj *model.RequestObject, clientObj clientmgr.IClient) *model.ResponseObject {
	var responseObj = model.NewResponseObject()
	var paramObj = new(LoginParameter)
	var rs model.ResponseStatus

	rs = requestObj.ParseParameter(&paramObj)
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	rs = paramObj.verify()
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	// Get existed player, or create a new player
	playerObj, exists := GetPlayerByName(paramObj.Name)
	if exists {
		// We should notify the current player that he/she is logging from a new device
		// Here I just skip this process
	} else {
		playerObj = playerModel.NewPlayer(paramObj.Name)
		register(playerObj)
	}

	// Assign a new room for this player
	roomObj, exists := room.AssignRoom()
	if !exists {
		return responseObj.SetResponseStatus(model.NoAvailableRoom)
	}
	roomObj.JoinRoom(playerObj)

	// Bind clientObj and playerObj together
	clientmgr.BindClientAndPlayer(clientObj, playerObj)

	// Return room's last 50 message
	responseObj.SetData(roomObj.GetMessageHistory())

	return responseObj
}

type StatsParameter struct {
	Name      string
	playerObj *playerModel.Player
}

func (this *StatsParameter) verify() model.ResponseStatus {
	if this.Name == "" {
		return model.ParamInValid
	}

	var exists bool
	this.playerObj, exists = GetPlayerByName(this.Name)
	if !exists {
		return model.PlayerNameNotExists
	}

	return model.Success
}

func stats(requestObj *model.RequestObject, clientObj clientmgr.IClient) *model.ResponseObject {
	var responseObj = model.NewResponseObject()
	var paramObj = new(StatsParameter)
	var rs model.ResponseStatus

	rs = requestObj.ParseParameter(&paramObj)
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	rs = paramObj.verify()
	if rs != model.Success {
		return responseObj.SetResponseStatus(rs)
	}

	// Get the required value
	statsStr := paramObj.playerObj.GetActiveTime()
	responseObj.SetData(statsStr)

	return responseObj
}
