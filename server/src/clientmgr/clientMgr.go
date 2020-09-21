/*clientMgr ...
In order to support multiple protocol simultaneously, such as socket, web socket, http and so on, I design this package.
serversocket provides support for tcp protocol.
serverwebsocket provides support for web socket protocol.
And clientMgr provides functions to manage all the client connections no matter what the underlying protocol is.
All the client connections must implement the IClient interface defined in interface.go
*/

package clientmgr

import (
	"sync"

	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

var (
	serverModuleName = "clientmgr"
	clientMap        = make(map[int64]IClient, 1024)
	mutex            sync.RWMutex
)

// RegisterClient ... Register a client into clientMap
func RegisterClient(clientObj IClient) {
	mutex.Lock()
	defer mutex.Unlock()

	clientMap[clientObj.GetID()] = clientObj
}

// UnregisterClient ... Unregister a client from clientMap
func UnregisterClient(clientObj IClient) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(clientMap, clientObj.GetID())
}

// GetClient ... Get client by id
// Return values:
// client object
// if exists or not
func GetClient(id int64) (clientObj IClient, exists bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	clientObj, exists = clientMap[id]
	return
}

func getClientCount() int {
	mutex.RLock()
	defer mutex.RUnlock()

	return len(clientMap)
}

func getClientList() (list []IClient) {
	mutex.RLock()
	defer mutex.RUnlock()

	for _, item := range clientMap {
		list = append(list, item)
	}
	return
}

func getExpiredClientList() (expiredList []IClient) {
	mutex.RLock()
	defer mutex.RUnlock()

	for _, item := range clientMap {
		if item.Expired() {
			expiredList = append(expiredList, item)
		}
	}

	return
}

func sendLoginAnotherDeviceMsg(clientObj IClient) {
	responseObj := model.NewResponseObject()
	responseObj.PushKey = model.LoginOnAnotherDevice

	clientObj.PlayerLogout()
	clientObj.SendMessage(responseObj)
	clientObj.SendMessage(model.NewDisconnectResponseObject())
}

// BindClientAndPlayer ... Bind client object and player object together after login
// clientObj: client object
// playerObj: player object
func BindClientAndPlayer(clientObj IClient, playerObj *playerModel.Player) {
	if playerObj.ClientID > 0 {
		if oldClientObj, exists := GetClient(playerObj.ClientID); exists {
			if clientObj != oldClientObj {
				sendLoginAnotherDeviceMsg(oldClientObj)
			}
		}
	}

	clientObj.PlayerLogin(playerObj)
	playerObj.ClientLogin(clientObj.GetID())
}

// Disconnect ... Disconnect a client from clientMap
// clientObj: client object
func Disconnect(clientObj IClient) {
	if playerObj := clientObj.GetPlayer(); playerObj != nil {
		playerLogoutHandler(playerObj)
	}

	clientObj.Close()
	UnregisterClient(clientObj)
}
