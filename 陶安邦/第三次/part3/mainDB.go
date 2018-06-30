package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"publicChain/part3/BLC"
)

func main() {




	//如果打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	////更新数据库
	//err = db.Update(func(tx *bolt.Tx) error {
	//
	//	//表名 bock
	//	//尝试读取表
	//	table := tx.Bucket([]byte("block"))
	//
	//	if table == nil {
	//		table,err= tx.CreateBucket([]byte("block"))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//	}
	//
	//
	//	genesisBlock := BLC.CreateGenesisBlock("Gennesis data...")
	//	blockByte := genesisBlock.Serialize()
	//
	//	table.Put([]byte("test"),blockByte)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//
	//	return nil
	//})



	//查看
	err = db.View(func(tx *bolt.Tx) error {

		table := tx.Bucket([]byte("block"))

		if table != nil {
			blockBytes :=  table.Get([]byte("test"))
			fmt.Println(blockBytes)

			block := BLC.DeserializeBlock(blockBytes)
			fmt.Println(block.Nonce)

		}


		return nil
	})


	//修改
	err = db.Batch(func(tx *bolt.Tx) error {

		return nil
	})



	}
