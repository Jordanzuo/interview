package model

// RequestObject ... Request object from client
type RequestObject struct {
	// Request id from client
	RequestID int64

	// Command
	Command string

	// Request data. This is a json object. Server will call function ParseParameter to parse this field into a specified object.
	// And the detailed definition is defined in each endpoint file.
	Parameter interface{}
}

func NewRequestObject(requestID int64, command string, parameter interface{}) *RequestObject {
	return &RequestObject{
		RequestID: requestID,
		Command:   command,
		Parameter: parameter,
	}
}
