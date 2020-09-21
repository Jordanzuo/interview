package serversocket

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Jordanzuo/goutil/intAndBytesUtil"
	"interview.com/cloudcade/chat/server/src/clientmgr"
	"interview.com/cloudcade/chat/server/src/model"
	playerModel "interview.com/cloudcade/chat/server/src/player/model"
)

const (
	conHeaderLength              = 4
	conClientExpireSeconds int64 = 300
)

var (
	byterOrder = binary.LittleEndian
)

// Client ... Client object, represents a socket connection
type Client struct {
	// Identifier
	id        int64
	requestId int64

	// underlying connection
	conn net.Conn

	// Received data
	receiveData []byte

	// Data waiting for sending to client
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

func (this *Client) getReceiveData() (message []byte, exists bool) {
	if len(this.receiveData) < conHeaderLength {
		return
	}
	header := this.receiveData[:conHeaderLength]
	contentLength := intAndBytesUtil.BytesToInt32(header, byterOrder)

	// If content length is zero, it represents heartbeat package.
	if contentLength == 0 {
		this.receiveData = this.receiveData[conHeaderLength:]
		if err := this.sendMessageToClient([]byte{}); err != nil {
			return
		}
		return
	}

	if len(this.receiveData) < conHeaderLength+int(contentLength) {
		return
	}

	// Extract message content
	message = this.receiveData[conHeaderLength : conHeaderLength+contentLength]
	exists = true
	this.receiveData = this.receiveData[conHeaderLength+contentLength:]

	return
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
	header := intAndBytesUtil.Int32ToBytes(int32(contentLength), byterOrder)
	message := append(header, content...)

	if err := this.sendMessageToClient(message); err != nil {
		return err
	}

	return nil
}

func (this *Client) sendMessageToClient(message []byte) error {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if _, err := this.conn.Write(message); err != nil {
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

		readBytes := make([]byte, 1024)
		n, err := this.conn.Read(readBytes)
		if err != nil {
			break
		}

		this.Active()

		this.receiveData = append(this.receiveData, readBytes[:n]...)
		atomic.StoreInt64(&this.activeTime, time.Now().Unix())

		for {
			message, exists := this.getReceiveData()
			if !exists {
				break
			}

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

func newClient(conn net.Conn) *Client {
	return &Client{
		id:          clientmgr.GetNewClientID(),
		conn:        conn,
		receiveData: make([]byte, 0, 1024),
		sendData:    make([]*model.ResponseObject, 0, 16),
		activeTime:  time.Now().Unix(),
		playerObj:   nil,
	}
}
