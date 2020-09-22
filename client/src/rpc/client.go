package rpc

import (
	"bytes"
	"encoding/binary"
	"net"
)

// BytesToInt32 ... Convert a byte arrty to an int32 value
func BytesToInt32(b []byte, order binary.ByteOrder) int32 {
	bytesBuffer := bytes.NewBuffer(b)
	var result int32
	binary.Read(bytesBuffer, order, &result)

	return result
}

// Int32ToBytes ... Convert an int32 value to byte array.
func Int32ToBytes(n int32, order binary.ByteOrder) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, order, n)
	return bytesBuffer.Bytes()
}

const (
	conHeaderLength = 4
)

var (
	byterOrder = binary.LittleEndian
)

type client struct {
	conn    net.Conn
	content []byte
}

func (clientObj *client) appendContent(content []byte) {
	clientObj.content = append(clientObj.content, content...)
}

func (clientObj *client) getValidMessage() ([]byte, bool) {
	if len(clientObj.content) < conHeaderLength {
		return nil, false
	}

	header := clientObj.content[:conHeaderLength]
	contentLength := BytesToInt32(header, byterOrder)
	if len(clientObj.content) < conHeaderLength+int(contentLength) {
		return nil, false
	}

	content := clientObj.content[conHeaderLength : conHeaderLength+contentLength]
	clientObj.content = clientObj.content[conHeaderLength+contentLength:]
	if contentLength == 0 || len(content) == 0 {
		return nil, false
	}

	return content, true
}

func (clientObj *client) sendByteMessage(message []byte) {
	contentLength := len(message)
	header := Int32ToBytes(int32(contentLength), byterOrder)
	message = append(header, message...)
	clientObj.conn.Write(message)
}

func (clientObj *client) sendStringMessage(s string) {
	clientObj.sendByteMessage([]byte(s))
}

func (clientObj *client) sendHeartBeatMessage() {
	clientObj.sendByteMessage([]byte{})
}

func newClient(_conn net.Conn) *client {
	return &client{
		conn:    _conn,
		content: make([]byte, 0, 1024),
	}
}
