package room

var (
	serverModuleName = "room"
	roomList         []*Room
)

// Init ... Init specified number of room. This can prevent data race condition related room.
func Init(roomCount int) {
	roomList = make([]*Room, roomCount, roomCount)
	for i := 0; i < roomCount; i++ {
		roomList[i] = newRoom(i + 1)
	}
}

func getRoomCount() int {
	return len(roomList)
}

func getRoom(id int) (roomObj *Room, exists bool) {
	if id < 1 || id > len(roomList) {
		return
	}

	roomObj = roomList[id-1]
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
