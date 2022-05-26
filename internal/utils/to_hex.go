package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ToHex(num int64) []byte {
	buff := bytes.NewBuffer([]byte{})
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panicln(err)
	}
	return buff.Bytes()
}
