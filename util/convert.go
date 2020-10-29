package util

import (
	"bytes"
	"encoding/binary"
)

//int---[]byte
func IntToBytes(num int64) ([]byte,error) {
	//bytes缓冲区
	buff := new(bytes.Buffer)
	//大端位序
	//小端位序
	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}

//int64---[]byte
//func Int64ToBytes(num int64) []byte {
//	return nil
//}

//string---[]byte
func StringToBytes(s string) []byte {
	return []byte(s)
}