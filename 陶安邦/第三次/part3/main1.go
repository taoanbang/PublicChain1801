package main

import (
	"publicChain/part3/BLC"
	"fmt"
)

func main() {
	//block := BLC.NewBlock("Genenis Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//fmt.Println(block)

	//genesisBlock := BLC.CreateGenesisBlock("Genenis block....")
	//fmt.Println(genesisBlock)

	//创建区块
	genesisBlockchain :=BLC.CreateBlockChainWithGenesisBlock()
	fmt.Println(genesisBlockchain.Blocks[0].Nonce)

	serialize := genesisBlockchain.Blocks[0].Serialize()

	fmt.Println(serialize)
	fmt.Println(BLC.DeserializeBlock(serialize).Nonce)

	genesisBlockchain.AddBlockToBlockchain(" send 200rmb to B",genesisBlockchain.Blocks[len(genesisBlockchain.Blocks)-1].Height,genesisBlockchain.Blocks[len(genesisBlockchain.Blocks)-1].Hash)





	}
