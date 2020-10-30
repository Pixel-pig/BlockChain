//将其他类型转为[]byte

package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// int64 转为 []byte
func Int64ToBytes(i int64) []byte {
	// 1.开辟一个 buff 缓存区
	buf := bytes.NewBuffer([]byte{})
	// 2.大端写入数据
	err := binary.Write(buf, binary.BigEndian, i)
	if err != nil {
		fmt.Println("int64 转为 []byte 错误")
		return nil
	}
	// 3.导出数据
	return buf.Bytes()
}

// string 转为 []byte
func StringToBytes(s string) []byte {
	return []byte(s)
}
