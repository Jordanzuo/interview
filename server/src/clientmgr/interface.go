package clientmgr

import (
	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

// IClient ... A interface for all types of clients. Such as socket and web socket.
type IClient interface {
	// Get client's id
	GetID() int64

	// Get client's request id
	GetClientRequestID() int64

	// Gssign a new request id to client
	SetClientRequestID(int64)

	// Get the player object binding to client
	GetPlayer() *playerModel.Player

	// Bind a player object to client
	PlayerLogin(*playerModel.Player)

	// Unbind player with client
	PlayerLogout()

	// Send message to client
	SendMessage(*model.ResponseObject)

	// Update active time
	Active()

	// Check if the client is expired
	Expired() bool

	// Close the underlying connection
	Close()

	// Quit the receive and send message loop
	Quit()
}
