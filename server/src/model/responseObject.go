package model

import (
	"time"
)

// ResponseObject ... Server's response object
type ResponseObject struct {
	// Response status to client
	ResponseStatus

	// Reponse data
	Data interface{}

	// Client's request distinct id
	RequestID int64

	// Timestamp
	TimeTick int64

	// Push key, used by server to push messages to client actively
	PushKey string
}

// IsDisconnect ... If this object represents a signal to disconnect the client connection
func (this *ResponseObject) IsDisconnect() bool {
	return this.ResponseStatus == DisconnectStatus
}

// SetResponseStatus ... Set a new response status value
// Return values:
// The same object
func (this *ResponseObject) SetResponseStatus(rs ResponseStatus) *ResponseObject {
	this.ResponseStatus = rs
	return this
}

// SetPushKey ... Set a new push key value
// Return values:
// The same object
func (this *ResponseObject) SetPushKey(pushKey string) *ResponseObject {
	this.PushKey = pushKey
	return this
}

// SetData ... Set a data value
// Return values:
// The same object
func (this *ResponseObject) SetData(data interface{}) *ResponseObject {
	this.Data = data
	return this
}

// SetTimeTick ... Set a timestamp
// Return values:
// The same object
func (this *ResponseObject) SetTimeTick(timeTick int64) *ResponseObject {
	this.TimeTick = timeTick
	return this
}

// NewResponseObject ... Create a new response object
// Return values:
// a new response object
func NewResponseObject() *ResponseObject {
	return &ResponseObject{
		ResponseStatus: Success,
		TimeTick:       time.Now().Unix(),
	}
}

// NewDisconnectResponseObject ... A short hand way to create a new response object and set the ResponseStatus field to DisconnectStatus
// Return values:
// a new response object
func NewDisconnectResponseObject() *ResponseObject {
	return &ResponseObject{
		ResponseStatus: DisconnectStatus,
	}
}
