package BLC

import (
	"math/big"
	"crypto/sha256"
	"fmt"
	"bytes"
)

//256位Hash里面前面至少要有16个零
const targetBit = 20

type ProofOfWork struct {
	Block *Block //当前要验证的区块
	target *big.Int //大数据存储
}

func NewProofOfWork(block *Block) *ProofOfWork{
	//创建一个初始化为1的target

	//int
	target := big.NewInt(1);

	// 左移256 - targetBit
	target = target.Lsh(target,256-targetBit)
	return &ProofOfWork{block,target}

}


func (ProofOfWork *ProofOfWork) Run()([]byte,int64){

	//讲Block的属性拼接成字节数组

	//2.生成hash

	//判断hash有效，如果满足条件 跳出循环

	nonce :=0

	var hashInt big.Int //存储我们新生成的hash
	var hash[32]byte

	for  {
		//准备数据
		dataBytes :=ProofOfWork.prepareData(nonce)

		//生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x",hash)

		//将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//判断hashInt是否小于Blcok里面的target
		if ProofOfWork.target.Cmp(&hashInt) == 1{
			break
		}

		nonce = nonce + 1


	}

	return hash[:],int64(nonce)

}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(nonce)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},

	)
	return data
}

