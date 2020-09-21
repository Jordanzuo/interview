package model

import (
	"encoding/json"
	"fmt"
)

// RequestObject ... Request object from client
type RequestObject struct {
	// Request id from client
	RequestID int64

	// Command
	Command string

	// Request data
	Parameter interface{}
}

// ParseParameter: Parse parameter to a specified object
func (this *RequestObject) ParseParameter(obj interface{}) ResponseStatus {
	bytes, err := json.Marshal(this.Parameter)
	if err != nil {
		fmt.Printf("ParseParameter.Marshal failed. Error:%v\n", err)
		return ClientDataError
	}

	err = json.Unmarshal(bytes, obj)
	if err != nil {
		fmt.Printf("ParseParameter.Unmarshal failed. Error:%v\n", err)
		return ClientDataError
	}

	return Success
}
