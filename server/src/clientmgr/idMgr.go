package clientmgr

import "sync/atomic"

var (
	globalClientID int64 = 0
)

// GetNewClientID ... Get a new client id for a new connection
func GetNewClientID() int64 {
	atomic.AddInt64(&globalClientID, 1)
	return globalClientID
}
