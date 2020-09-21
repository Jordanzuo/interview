package player

import (
	"sync"

	"interview.com/cloudcade/chat/server/src/clientmgr"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

func init() {
	clientmgr.RegisterPlayerLogoutHandler(clientLogout)
}

var (
	playerMap = make(map[int]*playerModel.Player, 1024)
	rwmutex   sync.RWMutex
)

func register(playerObj *playerModel.Player) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	playerMap[playerObj.ID] = playerObj
}

func unregister(playerObj *playerModel.Player) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	delete(playerMap, playerObj.ID)
}

func GetPlayerByID(id int) (playerObj *playerModel.Player, exists bool) {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	playerObj, exists = playerMap[id]
	return
}

func GetPlayerByName(name string) (playerObj *playerModel.Player, exists bool) {
	// This function can be optimized if this function is called very frequently.
	// The way to optimize this function is to store player in a map by using name as the key
	// In this way, it will reduce the response time but use more memory.
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	for _, v := range playerMap {
		if v.Name == name {
			playerObj = v
			exists = true
			return
		}
	}

	return
}

func getPlayerCount() int {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return len(playerMap)
}

func clientLogout(playerObj *playerModel.Player) {
	playerObj.ClientLogout()
	unregister(playerObj)
}
