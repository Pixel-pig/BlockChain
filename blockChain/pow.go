//挖矿的 Pow算法（工作量证明算法）

package blockChain

import (
	"blockChainProject/util"
	"bytes"
	"math/big"
)

const DIFFICULTY = 16

type POW struct {
	Target *big.Int
	Block  Block
}

//实例化pow算法对象
func NewPow(block Block) POW {
	target := big.NewInt(1)            //设置初始值
	target.Lsh(target, 255-DIFFICULTY) //位左移
	pow := POW{
		Target: target,
		Block:  block,
	}
	return pow
}

//执行pow算法
func (p POW) run() ([]byte, int64) {
	var nonce int64
	var blockHash []byte //整个区块的hash值
	bigBlock := new(big.Int)
	for {
		block := p.Block
		heightBytes := util.Int64ToBytes(block.Height)
		timeStampByte := util.Int64ToBytes(block.TimeStamp)
		versionByte := util.StringToBytes(block.Version)
		nonceByte := util.Int64ToBytes(nonce)
		blockBytes := bytes.Join([][]byte{
			heightBytes,
			timeStampByte,
			versionByte,
			nonceByte,
			block.Data,
			block.PrevHash,
		}, []byte{})
		//得到A + nonce 的值为[]byte
		blockHash = util.Sha256ToByte(blockBytes)
		//将byte转为big.Int
		bigBlock = bigBlock.SetBytes(blockHash)
		//判断大小
		if p.Target.Cmp(bigBlock) == 1 { //满足条件
			break
		}
		nonce++
	}
	return blockHash, nonce
}
