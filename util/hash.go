//hash 运算bao

package util

import (
	"crypto/sha256"
)

//// 1.将所有字段转为[]byte类型
//heightBytes, _ := Int64ToBytes(bk.Height)
//timeStampBytes, _ := Int64ToBytes(bk.TimeStamp)
//versionBytes := StringToBytes(bk.Version)
//// 2.拼接所有字段
//blockBytes := bytes.Join([][]byte{
//heightBytes,
//timeStampBytes,
//versionBytes,
//bk.Data,
//}, []byte{})

func Sha256ToByte(blockBytes []byte) []byte {
	// 1.将所有字段转为[]byte类型
	// 2.拼接所有字段
	// 1和2 抛出

	// 3.将拼接后的字段，hash
	sha := sha256.New()
	sha.Write(blockBytes)
	blockHash := sha.Sum(nil)
	return blockHash
}
