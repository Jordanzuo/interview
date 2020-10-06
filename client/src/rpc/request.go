package rpc

import (
	"encoding/json"
	"sync"
	"sync/atomic"

	"interview.com/cloudcade/chat/client/src/model"
)

var (
	globalRequestID     int64 = 0
	requestCallbackMap        = make(map[int64]func(interface{}), 64)
	requetCallbackMutex sync.Mutex
)

func generateNewRequestID() int64 {
	atomic.AddInt64(&globalRequestID, 1)
	return globalRequestID
}

func request(requestObj *model.RequestObject, callback func(interface{})) {
	message, _ := json.Marshal(requestObj)
	clientObj.sendByteMessage(message)

	requetCallbackMutex.Lock()
	defer requetCallbackMutex.Unlock()
	requestCallbackMap[requestObj.RequestID] = callback
}

func sendHeartBeat() {
	clientObj.sendByteMessage([]byte{})
}

func getCallback(requestID int64) (callback func(interface{}), exists bool) {
	requetCallbackMutex.Lock()
	defer requetCallbackMutex.Unlock()
	callback, exists = requestCallbackMap[requestID]
	delete(requestCallbackMap, requestID)
	return
}
