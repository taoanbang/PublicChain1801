package BLC

type Blockchain struct {
	Blocks []*Block  //存储有序的区块

}


func (blc *Blockchain) AddBlockToBlockchain(data string,height int64,preHash []byte){
	newBlock := NewBlock(data,height,preHash)
	blc.Blocks = append(blc.Blocks,newBlock)

}

//1.创建带有创世区块的区块链
func CreateBlockChainWithGenesisBlock() *Blockchain {
	genesisBlock := CreateGenesisBlock( "Genesis data...")
	return &Blockchain{[]*Block{genesisBlock}}

}




