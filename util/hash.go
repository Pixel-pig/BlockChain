//hash 运算包

package util

import (
	"crypto/sha256"
)

func Sha256ToByte(blockBytes []byte) []byte {
	sha := sha256.New()
	sha.Write(blockBytes)
	blockHash := sha.Sum(nil)
	return blockHash
}
