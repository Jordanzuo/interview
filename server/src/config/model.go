package config

// Config ... system config object
type Config struct {
	// The address that socket server will listen onn
	SocketServerListenAddr string

	// The address that web socket server will listen on
	WebSocketServerListenAddr string

	// The amount of rooms that the server will create and maintain
	RoomCount int
}
