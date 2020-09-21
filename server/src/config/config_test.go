package config

import (
	"testing"
)

func Test(t *testing.T) {
	err := Init("../../config/config.json")
	if err != nil {
		panic(err)
	}

	configObj := GetConfig()
	if configObj == nil {
		t.Errorf("There should be an non-empty config object. But now there is.")
		return
	}

	expectedSocketServerListenAddr := ":10001"
	gotSocketServerListenAddr := configObj.SocketServerListenAddr
	if expectedSocketServerListenAddr != gotSocketServerListenAddr {
		t.Errorf("SocketServerListenAddr: Expected to get %s, but now got %s.", expectedSocketServerListenAddr, gotSocketServerListenAddr)
		return
	}

	expectedWebSocketServerListenAddr := ":10002"
	gotWebSocketServerListenAddr := configObj.WebSocketServerListenAddr
	if expectedWebSocketServerListenAddr != gotWebSocketServerListenAddr {
		t.Errorf("WebSocketServerListenAddr: Expected to get %s, but now got %s.", expectedWebSocketServerListenAddr, gotWebSocketServerListenAddr)
		return
	}

	expectedRoomCount := 100
	gotRoomCount := configObj.RoomCount
	if expectedRoomCount != gotRoomCount {
		t.Errorf("RoomCount: Expected to get %d, but got %d.", expectedRoomCount, gotRoomCount)
		return
	}
}
