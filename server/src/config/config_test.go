package config

import (
	"testing"
)

func Test(t *testing.T) {
	err := Init("../../config/config.json")
	if err != nil {
		panic(err)
	}

	configObj := getConfig()
	if configObj == nil {
		t.Errorf("There should be an non-empty config object. But now there is.")
		return
	}

	expectedListenAddr := ":10001"
	gotListenAddr := configObj.ListenAddr
	if expectedListenAddr != gotListenAddr {
		t.Errorf("ListenAddr: Expected to get %s, but now got %s.", expectedListenAddr, gotListenAddr)
		return
	}

	expectedRoomCount := 100
	gotRoomCount := configObj.RoomCount
	if expectedRoomCount != gotRoomCount {
		t.Errorf("RoomCount: Expected to get %d, but got %d.", expectedRoomCount, gotRoomCount)
		return
	}
}
