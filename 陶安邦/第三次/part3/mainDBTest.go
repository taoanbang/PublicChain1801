package main

import (
	"publicChain/part3/BLC"
	"fmt"
	"encoding/hex"
)

func main() {

	blockchain2 := BLC.CreateBlockChainWithGenesisBlock()

	fmt.Println(hex.EncodeToString(blockchain2.Tip))

	//
	//blockchain.AddBlockToBlockchain("set1")
	//blockchain.AddBlockToBlockchain("set2")
	//
	//blockchain.AddBlockToBlockchain("set3")
	//blockchain.AddBlockToBlockchain("set4")
	//
	//fmt.Println(hex.EncodeToString(blockchain.Tip))



}
