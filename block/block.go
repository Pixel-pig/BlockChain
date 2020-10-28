//区块
//包含区块的基本字段，和相应的算法

package block

import (
	"BoltProject/util"
	"time"
)

// 区块的版本号
const (
	VERSION = "0x01"
)

type Block struct {
	Height    int64  // 区块高度
	TimeStamp int64  // 时间戳
	Hash      []byte //区块hash
	Data      []byte //数据
	PrevHash  []byte //上一个区块哈希
	Version   string // 版本号
	Nonce     int64  // 随机数 nonce
}

/**
 * 实例化一个区块，并赋值
 * Data数据在前端获取
 * PrevHash由之前的区块定义（创世区块除外）
 * Nonce Pow算法生成
 */
func (b Block) NewBlock(pHash []byte) Block {
	//实例化 block
	block := Block{
		Height:    b.Height + 1,
		TimeStamp: time.Now().Unix(),
		PrevHash:  pHash,
		Version:   VERSION,
	}

	return block
}

//创世区块的height为0，时间戳创造的时候获取，区块hash通过hash运算和pow获取，data数据有外部获取，phash为32个0，Nonce从pow算法获取
//data数据暂定位 （This block）
func (b Block) NewGenesisBlock() Block {
	data := util.StringToBytes("This block")
	//实例化创世区块
	block := Block{
		Height:    0,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Version:   VERSION,
	}
	return block
}
