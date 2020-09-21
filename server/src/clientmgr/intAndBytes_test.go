package clientmgr

import (
	"encoding/binary"
	"testing"
)

func TestInt32ToBytes(t *testing.T) {
	var expectedBigEndian []byte = []byte{0, 0, 1, 0}
	var expectedLittleEndian []byte = []byte{0, 1, 0, 0}
	var givenInt int32 = 256

	result := Int32ToBytes(givenInt, binary.BigEndian)
	if equal(result, expectedBigEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedBigEndian)
	}

	result = Int32ToBytes(givenInt, binary.LittleEndian)
	if equal(result, expectedLittleEndian) == false {
		t.Errorf("IntToBytes(%v) failed.Got %v, expected %v", givenInt, result, expectedLittleEndian)
	}
}
func TestBytesToInt32(t *testing.T) {
	var givenBigEndian []byte = []byte{0, 0, 1, 0}
	var givenLittleEndian []byte = []byte{0, 1, 0, 0}
	var expectedInt int32 = 256

	result := BytesToInt32(givenBigEndian, binary.BigEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenBigEndian, result, expectedInt)
	}

	result = BytesToInt32(givenLittleEndian, binary.LittleEndian)
	if result != expectedInt {
		t.Errorf("BytesToInt(%v) failed.Got %v, expected %v", givenLittleEndian, result, expectedInt)
	}
}
func equal(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}

	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}

	return true
}
