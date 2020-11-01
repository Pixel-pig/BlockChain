package main

import (
	"blockChainProject/blockChain"
	"fmt"
)

func main() {
	//测试区块链
	bc, err := blockChain.NewBlockChain()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(bc)
	//存储一个区块
	//block1 := blockChain.NewBlock(2, []byte("这是第三个区块"), bc.LastHash)
	//_, err := bc.SaveBlock(block1)
	//if err != nil {
	//	return
	//}
}
