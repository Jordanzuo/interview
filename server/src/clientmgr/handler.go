package clientmgr

import (
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

var (
	playerLogoutHandler func(*playerModel.Player)
)

// RegisterPlayerLogoutHandler ...
func RegisterPlayerLogoutHandler(handler func(*playerModel.Player)) {
	playerLogoutHandler = handler
}
