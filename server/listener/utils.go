package listener

import "encoding/binary"

func intToBytes(n int) []byte {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, uint32(n))
	return buff
}

func bytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}
