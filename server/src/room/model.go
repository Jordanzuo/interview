package room

import (
	"sync"

	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

const (
	conMaxPlayerCountPerRoom  = 100
	conMaxMessageCountPerRoom = 100
	conMaxMessageHistoryCount = 50
)

// Message ... The message in a room
type Message struct {
	PlayerName string
	Message    string
}

func newRoomMessage(playerName, message string) *Message {
	return &Message{
		PlayerName: playerName,
		Message:    message,
	}
}

// Room ... room object
type Room struct {
	// Identifier
	ID int

	// All the players in this room
	playerMap     map[int]*playerModel.Player
	playerRWMutex sync.RWMutex

	// All the message history
	messageList    []*Message
	messageRWMutex sync.RWMutex
}

func (this *Room) JoinRoom(playerObj *playerModel.Player) {
	this.playerRWMutex.Lock()
	defer this.playerRWMutex.Unlock()
	this.playerMap[playerObj.ID] = playerObj
	playerObj.JoinRoom(this.ID)
}

func (this *Room) exitRoom(playerObj *playerModel.Player) {
	this.playerRWMutex.Lock()
	defer this.playerRWMutex.Unlock()
	delete(this.playerMap, playerObj.ID)
}

func (this *Room) GetAllPlayers() []*playerModel.Player {
	this.playerRWMutex.RLock()
	defer this.playerRWMutex.RUnlock()

	allPlayerList := make([]*playerModel.Player, 0, len(this.playerMap))
	for _, v := range this.playerMap {
		allPlayerList = append(allPlayerList, v)
	}

	return allPlayerList
}

func (this *Room) getPlayerCount() int {
	this.playerRWMutex.RLock()
	defer this.playerRWMutex.RUnlock()
	return len(this.playerMap)
}

func (this *Room) isFull() bool {
	return this.getPlayerCount() >= conMaxPlayerCountPerRoom
}

func (this *Room) AppendMessage(playerObj *playerModel.Player, message string) {
	messageObj := newRoomMessage(playerObj.Name, message)

	this.messageRWMutex.Lock()
	defer this.messageRWMutex.Unlock()
	this.messageList = append(this.messageList, messageObj)

	// Check if the message's count has exceeded a certain number?
	// If so, just leave the last con_Max_Message_History_Count messages.
	if len(this.messageList) > conMaxMessageCountPerRoom {
		this.messageList = this.messageList[len(this.messageList)-conMaxMessageHistoryCount:]
	}
}

func (this *Room) GetMessageHistory() []*Message {
	this.messageRWMutex.RLock()
	defer this.messageRWMutex.RUnlock()

	retMessageList := make([]*Message, 0, conMaxMessageHistoryCount)
	if len(this.messageList) < conMaxMessageHistoryCount {
		for _, v := range this.messageList {
			retMessageList = append(retMessageList, v)
		}
	} else {
		for _, v := range this.messageList[len(this.messageList)-conMaxMessageHistoryCount:] {
			retMessageList = append(retMessageList, v)
		}
	}

	return retMessageList
}

func newRoom(id int) *Room {
	return &Room{
		ID:          id,
		playerMap:   make(map[int]*playerModel.Player, 64),
		messageList: make([]*Message, 0, 64),
	}
}
