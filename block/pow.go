//挖矿的 Pow算法（工作量证明算法）

package block

import (
	"BoltProject/util"
	"bytes"
	"math/big"
)

type pow struct {
	Difficulty int
	Target big.Int
}

//
//func

/**
 * 1.拿到区块
 * 2.对区块信息进行拼接(数据信息全为[]byte类型[在进行hash运算时参数均为[]byte])
 * 3.对区块信息进行hash运算
 * 4.利用A + nonce < B获取最终的hash值（A + nonce）pow的核心 {比较时用大正数比较，因为其他类型的范围不足}
 * 5.赋值hash，导出block info
 */
func (p pow)run(block Block) Block {
	heightBytes := util.Int64ToBytes(block.Height)
	timeStampByte := util.Int64ToBytes(block.TimeStamp)
	versionByte := util.StringToBytes(block.Version)
	var nonce int64 = 0
	for {
		nonceByte := util.Int64ToBytes(nonce)
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeStampByte,
			versionByte,
			nonceByte,
			block.Data,
			block.PrevHash,
		}, []byte{})
		blockHash := util.Sha256ToByte(blockBytes)
		nonce++
	}
	return block
}
