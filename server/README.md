# Environment
go version go1.14.1 linux/amd64
git version 2.25.1
Visual Studio Code: 1.49.1 with vim extension

# Endpoints
All the endpoints are defined in the file endpoint.go under corresponding paths.

## RequestObject:
The definition of RequestObject is as follows:

type RequestObject struct {
	// Request id from client
	RequestID int64

	// Command
	Command string

    // Request data. This is a json object. Server will call function ParseParameter to parse this field into a specified object
    // And the detailed definition is defined in each endpoint file.
	Parameter interface{}
}

## ResponseObject:
The definition of ResponseObject is as follows:

// ResponseObject ... Server's response object
type ResponseObject struct {
	// Response status to client
	ResponseStatus

	// Reponse data. This is a object which is defined by each endpoint.
	Data interface{}

	// Client's request distinct id. Just send back to client.
	RequestID int64

	// Timestamp
	TimeTick int64

	// Push key, used by server to push messages to client actively
	PushKey string
}