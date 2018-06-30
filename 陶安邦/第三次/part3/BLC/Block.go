package BLC

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	//1.区块链高度
	Height int64

	//2.上一个区块hash
	PrevBlockHash []byte

	//3.交易数据
	Data []byte

	//4.时间戳
	Timestamp int64

	//5.Hash
	Hash []byte

	//6.nonce
	Nonce int64


}

//func (block *Block) SetHash() {
//	//1.Height
//	heightBytes := IntToHex(block.Height)
//	fmt.Println(heightBytes )
//
//	//2. 时间戳转 []byte
//
//	//2 -  36
//	timeString :=strconv.FormatInt(block.Timestamp,2)
//
//	fmt.Println(timeString)
//	timesBytes :=[]byte(timeString)
//	fmt.Println(timesBytes)
//
//	//3 拼接
//	blockBytes := bytes.Join([][]bye{heightBytes,block,PrevBlockHash,block})
//
//	//4.生成Hash
//	hash :=sha256.Sum256(blockBytes)
//	block.Hash = hash[:]
//
//	fmt.Println(string(block.Hash))
//	fmt.Println(string(block.Nonce))
//
//
//	return block
//
//}


//1.创建新的区块
func NewBlock(data string,height int64,prevBlockHash []byte) *Block {

	block := &Block{Height:height,PrevBlockHash:prevBlockHash,Data:[]byte(data),Timestamp:time.Now().Unix(),Hash:nil}
	pow :=NewProofOfWork(block)

	hash,nonce :=pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	//fmt.Println(string(block.Hash))

	////设置hash
	//block.SetHash()

	return block


}

//2.单独写一个方法，生成创世区块

func CreateGenesisBlock(data string) *Block {

	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}


//将区块序列化成字节数组
func(block *Block) Serialize() []byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()

}

//反序列化
func DeserializeBlock(blockBytes []byte) *Block{
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}