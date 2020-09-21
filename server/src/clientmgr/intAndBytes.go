package clientmgr

import (
	"bytes"
	"encoding/binary"
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
