package model

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	// The max seconds for a player before it's declared expired.
	con_Player_Expire_Seconds = 300
	con_Seconds_Per_Day       = 86400
	con_Seconds_Per_Hour      = 3600
	con_Seconds_Per_Minute    = 60
)

var (
	globalID int32
)

func getNewId() int {
	newID := atomic.AddInt32(&globalID, 1)
	return int(newID)
}

// Player ... player object
type Player struct {
	// Identifier
	ID int

	// Player's name
	Name string

	// Login time
	LoginTime int64

	// Player's last active time
	ActiveTime int64

	// The connection's id
	ClientID int64

	// The id of a room which this player belongs to
	RoomID int
}

func (this *Player) ClientLogin(clientID int64) {
	this.ClientID = clientID
	this.LoginTime = time.Now().Unix()
}

func (this *Player) ClientLogout() {
	this.ClientID = 0
	this.RoomID = 0
}

func (this *Player) JoinRoom(roomID int) {
	this.RoomID = roomID
	fmt.Printf("%s joins room:%d\n", this.Name, roomID)
}

func (this *Player) UpdateActiveTime(activeTime int64) {
	this.ActiveTime = activeTime
}

// Check if player is expired
func (this *Player) Expired() bool {
	if this.ClientID == 0 && time.Now().Unix()-this.ActiveTime > con_Player_Expire_Seconds {
		return true
	}

	return false
}

func (this *Player) GetActiveTime() string {
	totalActiveSeconds := time.Now().Unix() - this.LoginTime
	days, hours, minutes, seconds := int64(0), int64(0), int64(0), int64(0)
	days, totalActiveSeconds = totalActiveSeconds/con_Seconds_Per_Day, totalActiveSeconds%con_Seconds_Per_Day
	hours, totalActiveSeconds = totalActiveSeconds/con_Seconds_Per_Hour, totalActiveSeconds%con_Seconds_Per_Hour
	minutes, totalActiveSeconds = totalActiveSeconds/con_Seconds_Per_Minute, totalActiveSeconds%con_Seconds_Per_Minute
	seconds = totalActiveSeconds

	return fmt.Sprintf("%2dd %2dh %2dm %2ds", days, hours, minutes, seconds)
}

func NewPlayer(name string) *Player {
	return &Player{
		ID:   getNewId(),
		Name: name,
	}
}
