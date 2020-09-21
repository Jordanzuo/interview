package serverwebsocket

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"interview.com/cloudcade/chat/server/src/clientmgr"
	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

const (
	conMaxMessageSize      = 1024
	conClientExpireSeconds = 30
)

var (
	byterOrder = binary.LittleEndian
)

// Client ... Client object, represents a socket connection
type Client struct {
	// Identifier
	id        int64
	requestId int64

	// The websocket connection.
	conn *websocket.Conn

	// The websocket connection.
	sendData []*model.ResponseObject

	// Is closed
	closed     bool
	playerObj  *playerModel.Player
	activeTime int64
	mutex      sync.Mutex
}

func (this *Client) GetID() int64 {
	return this.id
}

func (this *Client) GetClientRequestID() int64 {
	return this.requestId
}

func (this *Client) SetClientRequestID(requestId int64) {
	this.requestId = requestId
}

func (this *Client) GetPlayer() *playerModel.Player {
	return this.playerObj
}

func (this *Client) PlayerLogin(playerObj *playerModel.Player) {
	this.playerObj = playerObj
}

func (this *Client) PlayerLogout() {
	this.playerObj = nil
}

func (this *Client) getSendData() (responseObj *model.ResponseObject, exists bool) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if len(this.sendData) == 0 {
		return
	}

	responseObj = this.sendData[0]
	exists = true
	this.sendData = this.sendData[1:]

	return
}

func (this *Client) SendMessage(responseObj *model.ResponseObject) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.sendData = append(this.sendData, responseObj)
}

func (this *Client) sendMessage(responseObj *model.ResponseObject) error {
	content, _ := json.Marshal(responseObj)
	contentLength := len(content)
	header := clientmgr.Int32ToBytes(int32(contentLength), byterOrder)
	message := append(header, content...)

	if err := this.sendMessageToClient(message); err != nil {
		return err
	}

	return nil
}

func (this *Client) sendMessageToClient(message []byte) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if err := this.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
		return err
	}

	return nil
}

func (this *Client) Active() {
	this.activeTime = time.Now().Unix()
	if this.playerObj != nil {
		this.playerObj.UpdateActiveTime(this.activeTime)
	}
}

func (this *Client) Expired() bool {
	return time.Now().Unix() > this.activeTime+conClientExpireSeconds
}

func (this *Client) Close() {
	this.conn.Close()
	this.closed = true
	fmt.Printf("ClientId:%d close\n", this.id)
}

func (this *Client) Quit() {
	this.conn.Close()
	this.closed = true
	clientmgr.UnregisterClient(this)
}

func (this *Client) start() {
	go this.handleReceiveData()
	go this.handleSendData()
}

func (this *Client) handleReceiveData() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer this.Quit()

	for {
		if this.closed {
			break
		}

		_, message, err := this.conn.ReadMessage()
		if err != nil {
			break
		}

		this.Active()

		if len(message) == 0 {
			err := this.sendMessageToClient([]byte{})
			if err != nil {
				break
			}
		} else {
			clientmgr.HandleRequest(this, message)
		}
	}
}

func (this *Client) handleSendData() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer this.Quit()

	for {
		if this.closed {
			break
		}

		sendDataItemObj, exists := this.getSendData()
		if exists {
			if sendDataItemObj.IsDisconnect() {
				break
			}

			err := this.sendMessage(sendDataItemObj)
			if err != nil {
				break
			}
		} else {
			time.Sleep(2 * time.Millisecond)
		}
	}
}

func newClient(conn *websocket.Conn) *Client {
	conn.SetReadLimit(conMaxMessageSize)

	return &Client{
		id:         clientmgr.GetNewClientID(),
		conn:       conn,
		sendData:   make([]*model.ResponseObject, 0, 16),
		activeTime: time.Now().Unix(),
		playerObj:  nil,
	}
}
