package model

// ResponseObject ... Server's response object
type ResponseObject struct {
	// Response status to client
	ResponseStatus string

	// Reponse data
	Data interface{}

	// Client's request distinct id
	RequestID int64

	// Timestamp
	TimeTick int64

	// Push key, used by server to push messages to client actively
	PushKey string
}
