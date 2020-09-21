package clientmgr

import (
	"fmt"
	"time"
)

func init() {
	go clearExpiredClient()
}

// clear expired client
func clearExpiredClient() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[%s]Goroutine: clearExpiredClient encounters some error: %v", serverModuleName, r)
		}
	}()

	for {
		// Sleep for a while, because there is no need to handle expired clients when the system is just starting up.
		time.Sleep(5 * time.Second)

		// Get expired client list and disconnect all of them
		beforeClientCount := getClientCount()
		expiredClientList := getExpiredClientList()
		expiredClientCount := len(expiredClientList)
		if expiredClientCount > 0 {
			for _, item := range expiredClientList {
				Disconnect(item)
			}

			fmt.Printf("[%s]The num of clients before cleaning：%d， cleaned num：%d\n", serverModuleName, beforeClientCount, expiredClientCount)
		}
	}
}
