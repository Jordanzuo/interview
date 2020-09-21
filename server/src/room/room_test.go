package room

import (
	"fmt"
	"testing"

	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

func TestRoomPlayer(t *testing.T) {
	id := 1
	roomObj := newRoom(id)

	expectedPlayerCount := 0
	allPlayerList := roomObj.GetAllPlayers()
	if len(allPlayerList) != expectedPlayerCount {
		t.Errorf("Expected to get %d players. But now there are %d players.", expectedPlayerCount, len(allPlayerList))
		return
	}

	playerList := make([]*playerModel.Player, 0, 100)
	for i := 1; i <= 100; i++ {
		playerObj := playerModel.NewPlayer(fmt.Sprintf("Player%d", i))
		playerList = append(playerList, playerObj)
		roomObj.JoinRoom(playerObj)
		expectedPlayerCount = i
		allPlayerList = roomObj.GetAllPlayers()
		if len(allPlayerList) != expectedPlayerCount {
			t.Errorf("Expected to get %d players. But now there are %d players.", expectedPlayerCount, len(allPlayerList))
			return
		}
	}

	for i := 1; i <= 100; i++ {
		playerObj := playerList[i-1]
		roomObj.ExitRoom(playerObj)
		expectedPlayerCount = 100 - i
		allPlayerList = roomObj.GetAllPlayers()
		if len(allPlayerList) != expectedPlayerCount {
			t.Errorf("Expected to get %d players. But now there are %d players.", expectedPlayerCount, len(allPlayerList))
			return
		}
	}
}

func TestRoomMessage(t *testing.T) {
	id := 1
	roomObj := newRoom(id)
	playerObj := playerModel.NewPlayer("Player1")

	expectedMessageCount := 0
	messageHistoryList := roomObj.GetMessageHistory()
	if len(messageHistoryList) != expectedMessageCount {
		t.Errorf("Expected to get %d messages. But now there are %d messages.", expectedMessageCount, len(messageHistoryList))
		return
	}

	for i := 1; i < 2*con_Max_Message_Count_Per_Room; i++ {
		roomObj.AppendMessage(playerObj, fmt.Sprintf("Test message_%d", i))
		if i <= con_Max_Message_History_Count {
			expectedMessageCount = i
		} else {
			expectedMessageCount = con_Max_Message_History_Count
		}

		messageHistoryList = roomObj.GetMessageHistory()
		if len(messageHistoryList) != expectedMessageCount {
			t.Errorf("Expected to get %d messages. But now there are %d messages.", expectedMessageCount, len(messageHistoryList))
			return
		}
	}
}

func TestAssignRoom(t *testing.T) {
	roomCount := 100
	Init(roomCount)

	expectedRoomCount := roomCount
	gotRoomCount := getRoomCount()
	if expectedRoomCount != gotRoomCount {
		t.Errorf("Expected to get %d rooms. But now there are %d rooms.", expectedRoomCount, gotRoomCount)
		return
	}

	for i := 0; i < 15000; i++ {
		playerObj := playerModel.NewPlayer(fmt.Sprintf("Player_%d", i))
		newRoomObj, exists := AssignRoom()
		if i >= roomCount*con_Max_Player_Count_Per_Room {
			if exists {
				t.Errorf("There should be no available room for new player. But now there is")
				return
			}
		} else {
			if !exists {
				t.Errorf("There should be enough available room for new player. But now there isn't.")
				return
			}

			expectedRoomID := i / con_Max_Player_Count_Per_Room
			if newRoomObj.ID != expectedRoomID {
				t.Errorf("Expected to get a room with ID:%d. But now it's %d", expectedRoomID, newRoomObj.ID)
				return
			}
			newRoomObj.JoinRoom(playerObj)
		}
	}
}
