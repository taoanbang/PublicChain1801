package main

import (
	"publicChain/part2-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	//block := BLC.NewBlock("Genenis Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//fmt.Println(block)

	//genesisBlock := BLC.CreateGenesisBlock("Genenis block....")
	//fmt.Println(genesisBlock)

	//创建区块
	genesisBlockchain :=BLC.CreateBlockChainWithGenesisBlock()
	fmt.Println(genesisBlockchain.Blocks)
	fmt.Println(genesisBlockchain.Blocks[0])

	genesisBlockchain.AddBlockToBlockchain(" send 200rmb to B",genesisBlockchain.Blocks[len(genesisBlockchain.Blocks)-1].Height,genesisBlockchain.Blocks[len(genesisBlockchain.Blocks)-1].Hash)





	}
