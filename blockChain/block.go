//区块
//包含区块的基本字段，和相应的算法

package blockChain

import (
	"bytes"
	"encoding/gob"
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

//创世区块的height为0，时间戳创造的时候获取，区块hash通过hash运算和pow获取，data数据有外部获取，phash为32个0，Nonce从pow算法获取
//data数据暂定位 （This block）
func NewGenesisBlock() Block {
	//实例化创世区块
	data := []byte("这是创世区块")
	genesisBlock := NewBlock(0, data, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return genesisBlock
}

//创建一个新的区块
func NewBlock(height int64, data []byte, prevHash []byte) Block {
	//1.实例化一个生成区块
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   VERSION,
	}
	//2.执行pow算法寻找nonce值
	blockHash, nonce := NewPow(block).run()
	block.Nonce = nonce
	block.Hash = blockHash

	return block
}

//gob 序列化
func (b Block) Serialize() ([]byte, error) {
	//1. 开辟缓冲区
	var buf bytes.Buffer
	//2. 创建编码器
	encoder := gob.NewEncoder(&buf)
	//3. 编码
	err := encoder.Encode(b)
	//4. 序列化后的值
	structInfo := buf.Bytes()
	if err != nil {
		return nil, err
	}
	return structInfo, nil
}

/**
 * gob 反序列化(b,)
 * 1. 获取序列化后的值
 * 2. 创建解码器
 * 3. 解码
 */
func DeSerialize(structInfo []byte) (*Block, error) {
	var block Block
	//创建解码器
	decoder := gob.NewDecoder(bytes.NewReader(structInfo))
	//解码
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}
