package rpc

import (
	"time"
)

func init() {
	go func() {
		for {
			time.Sleep(5 * time.Second)

			if clientObj == nil {
				continue
			}

			sendHeartBeat()
		}
	}()
}
