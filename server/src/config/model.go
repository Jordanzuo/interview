package config

// Config ... system config object
type Config struct {
	ListenAddr string // The address this server listens on
	RoomCount  int    // The room count that this server creates and maintains
}
