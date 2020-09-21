package room

import (
	"fmt"
	"time"
)

func init() {
	go clearOfflinePlayer()
}

// clear offline player
func clearOfflinePlayer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[%s]Goroutine: clearOfflinePlayer encounters some error: %v", serverModuleName, r)
		}
	}()

	for {
		// Sleep for a while, because there is no need to handle offline clients when the system is just starting up.
		time.Sleep(5 * time.Second)

		for _, roomObj := range roomList {
			for _, playerObj := range roomObj.GetAllPlayers() {
				if playerObj.RoomID == 0 {
					roomObj.exitRoom(playerObj)
				}
			}
		}
	}
}
