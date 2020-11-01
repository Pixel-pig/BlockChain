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
	fmt.Println(bc.LastHash)
	gb := bc.QuaryBlockByHeight(0)
	fmt.Println(string(gb.Data))
	blocks := bc.QuaryAllBlock()
	for i := 0; i <len(blocks); i++ {
		fmt.Println(blocks[i])
	}
}
