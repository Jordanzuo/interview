package player

import (
	"fmt"
	"testing"

	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

func TestPlayer(t *testing.T) {
	expectedPlayerCount := 0
	gotPlayerCount := getPlayerCount()
	if expectedPlayerCount != gotPlayerCount {
		t.Errorf("Expected to get %d players. But now there are %d.", expectedPlayerCount, gotPlayerCount)
		return
	}

	count := 100
	for i := 0; i < count; i++ {
		playerObj := playerModel.NewPlayer(fmt.Sprintf("Player_%d", i))
		register(playerObj)
		expectedPlayerCount = i + 1
		gotPlayerCount = getPlayerCount()
		if expectedPlayerCount != gotPlayerCount {
			t.Errorf("Expected to get %d players. But now there are %d.", expectedPlayerCount, gotPlayerCount)
			return
		}
	}

	for i := 0; i < count; i++ {
		playerObj, exists := GetPlayerByName(fmt.Sprintf("Player_%d", i))
		if !exists {
			t.Errorf("There should be a player named Player_%d, but now there isn't.", i)
			return
		}

		_, exists = GetPlayerByID(playerObj.ID)
		if !exists {
			t.Errorf("There should be a player named Player_%d, but now there isn't.", i)
			return
		}

		unregister(playerObj)
		expectedPlayerCount = count - i - 1
		gotPlayerCount = getPlayerCount()
		if expectedPlayerCount != gotPlayerCount {
			t.Errorf("Expected to get %d players. But now there are %d.", expectedPlayerCount, gotPlayerCount)
			return
		}
	}
}
