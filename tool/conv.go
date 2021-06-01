package tool

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

/**
 * @Author: Lee
 * @Date: 2021/6/1 14:44
 * @Desc:
 */
//整形转换成字节

func Int8ToByte(n int) []byte {
	x := int8(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
func Int16ToByte(n int) []byte {
	x := int16(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
func Int32ToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}


func Hex2int64(val string) (int64,error) {
	n, err := strconv.ParseUint(val, 16, 64)
	if err != nil {
		return 0,err
	}
	return int64(n),nil
}
func Hex2int32(val string) (int32,error) {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		return 0,err
	}
	return int32(n),nil
}