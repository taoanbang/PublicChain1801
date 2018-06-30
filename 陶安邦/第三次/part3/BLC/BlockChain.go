package BLC

import (

	"publicChain/part3/BLC"
	"github.com/boltdb/bolt"
	"log"

	"math/big"
)



//数据库名字
const dbName = "blockchain.db"
//表名字
const blockTableName = "blocks"

//type Blockchain struct {
//	Blocks []*Block  //存储有序的区块
//
//}


type Blockchain struct {
	Tip []byte //最新的区块链Hash
	DB *bolt.DB

}


//func (blc *Blockchain) AddBlockToBlockchain(data string,height int64,preHash []byte){
//	newBlock := NewBlock(data,height,preHash)
//	blc.Blocks = append(blc.Blocks,newBlock)
//
//}

//1.创建带有创世区块的区块链
//func CreateBlockChainWithGenesisBlock() *Blockchain {
//	genesisBlock := CreateGenesisBlock( "Genesis data...")
//	return &Blockchain{[]*Block{genesisBlock}}
//
//}

func CreateBlockChainWithGenesisBlock() *Blockchain {
	//创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}


	//var blockHash []byte

	var blockchain *Blockchain

	//更新数据库
	err = db.Update(func(tx *bolt.Tx) error {




			//创建数据库表
			b,err := tx.CreateBucket([]byte(blockTableName))
			if err != nil {
				log.Panic(err)
			}

			if b != nil  {
				//创建创世区块
				genesisBlock := CreateGenesisBlock("Gennesis data...")

				err := b.Put(genesisBlock.Hash,genesisBlock.Serialize())

				if err != nil {
					log.Panic(err)
				}

				//存储最新的区块的hash
				err = b.Put([]byte("l"),genesisBlock.Hash)
				if err != nil{
					log.Panic(err)
				}

				blockchain = &Blockchain{gensisBlock.Hash, db}


			}


		return nil

	})



	return blockchain


	}




func (blc *Blockchain) AddBlockToBlockchain(data string){

	err := blc.DB.Update(func(tx *bolt.Tx) error{
		//1.获取表
		b := tx.Bucket([]byte(blockTableName))
		//2.创建新区块
		if b != nil {

			//先获取最新区块
			blockBytes := b.Get(blc.Tip)
			//反序列
			block := DeserializeBlock(blockBytes)

			//3讲区块序列化并且存储到数据库中
			newBlock := NewBlock(data,block.Height + 1,block.Hash)
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if err != nil {
				log.Println(err)
			}
			//4更新
			 err = b.Put([]byte["l"],newBlock.Hash)
			 if err != nil{
			 	log.Println(err)
			 }

			 //5更新
			 blc.Tip = newBlock.Hash

		}

		return nil

	})

	if err != nil {
		log.Println(err)
	}

}

type BlockchainIterator struct{

	CurrentHash []byte
	DB *bolt.DB
}

func (blockchain *Blockchain) Iterator() *BlockchainIterator{

	return &BlockchainIterator{CurrentHash:blockchain.Tip, DB:blockchain.DB}

}

func (blcIterator *BlockchainIterator) Next() * Block{

	var block *Block
	err := blcIterator.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b != nil{
			blockByes := b.Get(blcIterator.CurrentHash)
			block = DeserializeBlock(blockByes)

			blcIterator.CurrentHash = block.PrevBlockHash


		}
		return nil
	})

	if err != nil{
		log.Panic(err)
	}

	return block

}


func (blc *Blockchain) PrintChain(){
	var block *Block
	blcIterator := blc.Iterator()

	for{
		block := blcIterator.Next()
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)
		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break;
		}


	}


}


