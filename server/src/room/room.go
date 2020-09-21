package room

var (
	roomList []*Room
)

// Init ... Init specified number of room. This can prevent data race condition related room.
func Init(roomCount int) {
	roomList = make([]*Room, roomCount, roomCount)
	for i := 0; i < roomCount; i++ {
		roomList[i] = newRoom(i)
	}
}

func getRoomCount() int {
	return len(roomList)
}

func getRoom(id int) (roomObj *Room, exists bool) {
	if id < 0 || id > len(roomList)-1 {
		return
	}

	roomObj = roomList[id]
	exists = true

	return
}

// AssignRoom ... Assign a room for a newly coming player.
// Return values:
// a candidate room
// if there is an available room
func AssignRoom() (roomObj *Room, exists bool) {
	for _, v := range roomList {
		if v.isFull() == false {
			roomObj = v
			exists = true
			return
		}
	}

	return
}
