package main

import (
	"blockChainProject/blockChain"
)

func main() {
	//测试区块链
	blockChain.NewBlockChain()
	//存储一个区块
	//block1 := blockChain.NewBlock(2, []byte("这是第三个区块"), bc.LastHash)
	//_, err := bc.SaveBlock(block1)
	//if err != nil {
	//	return
	//}
}
