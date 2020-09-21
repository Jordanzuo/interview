package model

import (
	"time"
)

// ResponseObject ... 服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值
	ResponseStatus

	// 响应结果的数据
	Data interface{}

	// 客户端请求唯一标识
	RequestID int64

	// 时间戳
	TimeTick int64

	// 推送Key,用于主动推送的数据
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
