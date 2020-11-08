package blockchain

import (
	"DataCertProject/models"
	"errors"
	"fmt"
	"github.com/bolt-master"
	"math/big"
)

//桶的名称
var BUCKET_NAME = "blocks"
//表示最新的区块lashHash的key的名字
var LAST_KEY = "lasthash"
//存储区块数据的文件
var CHAINDB = "chain.db"

var CHAIN BlockChain
/**
*	区块链结构体实际定义：用于表示代表一条链
	该区块链包含以下功能：
		① 将新产生的区块与已有的区块链连起来，并保存
		② 可以查询某个区块的信息
		③ 可以将所有区块进行遍历，输出区块信息
*/

type BlockChain struct {
	LastHash []byte //最后一个hash
	BoltDb *bolt.DB
}

/**
 *
 */
func (bc BlockChain) QueryAllBlocks() []*Block {
	blocks := make([]*Block, 0)

	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据出错")
		}
		eachKey := bc.LastHash
		preHashBig := new(big.Int)
		zeroBig := big.NewInt(0)// 0 的大整数
		for  {
			eachBlockBytes := bucket.Get(eachKey)
			eachBlock, _ := DeSerialize(eachBlockBytes)//反序例化
			//遍历把每个区块结构体指针放入到[]byte容器中
			blocks = append(blocks,eachBlock)

			preHashBig.SetBytes(eachBlock.PrevHash)
			if preHashBig.Cmp(zeroBig) == 0 {//判断是否遍历完，是否到了创世区块，到了就停止
				break
			}//否则继续向前遍历

			eachKey = eachBlock.PrevHash
		}
		return nil
	})
	return blocks
}

/**
*	通过区块的高度查询某个具体的区块，返回区块实例
*/
func (bc BlockChain) QueryBlockByHeight(height int64) (*Block) {
	if height < 0 {//如果要查询的高度小于0，直接结束
		return nil
	}
	var block *Block
	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据失败")
		}
		hashKey := bc.LastHash
		for  {
			lastBlockBytes := bucket.Get(hashKey)
			eachBlock, _ := DeSerialize(lastBlockBytes)
			if eachBlock.Height < height {//判断查询的数据区块中是否存在，对比高度
				break
			}
			if eachBlock.Height == height {//高度和目标一致，结束循环
				block = eachBlock
				break
			}
			//遍历的当前的区块高度与目标高度不一致，继续往前遍历
			//以eachBLock.PrevHash为key，使用Get获取上一个区块的数据
			hashKey = eachBlock.PrevHash
		}

		return nil
	})

	return block
}

/**
  *用于创建一条区块链，并返回该区块链实例
	区块链是链，所以得从创世区块开始
  */
func NewBlockChain() BlockChain {
	//先打开文件
	db, err :=bolt.Open(CHAINDB,0600,nil)
	if err != nil {
		panic(err.Error())
	}

	var bl BlockChain
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 {//无创世区块时，也就是第一次构造区块链
			//创建第一个区块
			genesis := CreateGenesisBlock()//创世区块

			fmt.Printf("genesis的hash值：%x\n",genesis.Hash)

			bl = BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
			//将创世区块存入桶中，作为一条区块链的开端
			genesieBytes, err := genesis.Serialze()
			if err != nil {
				panic(err.Error())
			}
			bucket.Put(genesis.Hash,genesieBytes)
			bucket.Put([]byte(LAST_KEY),genesis.Hash)
			bl.LastHash = genesis.Hash
		}else {//又创世区块，就是再次往链里面存区块时
			lastHash := bucket.Get([]byte(LAST_KEY))
			lastBlockBytes := bucket.Get(lastHash)
			lastBlock, err := DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取数据失败")
			}
			
			bl = BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil

	})
	//为全局变量赋值
	CHAIN = bl

	return bl
}

/*
 *根据用户传入的保全号查询区块的信息，并返回
 */

func (bc BlockChain) QueryBlockByCertId(cert_id []byte) (*Block, error) {
	var block *Block
	db := bc.BoltDb
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("查询区块数据遇到错误！")
		}
		//未遇到错误，则桶存在
		eachHash := bucket.Get([]byte(LAST_KEY))
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)
		var certRecord *models.CertRecord
		for {
			eachBlockBytes := bucket.Get(eachHash)
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//如果数据存在，找得到,反序列化后对比
			certRecord, _ = models.DeSerializeRecord(eachBlock.Data)
			fmt.Println(string(certRecord.CertId))
			fmt.Println(string(cert_id))
			if string(certRecord.CertId) == string(cert_id) {
				block = eachBlock
				break
			}
			//如果数据不存在，找不到数据
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0 {//一直到创世区块都未找到
				break
			}

			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return block, err
}

//调用BlockChain的该SaveBlock方法，该方法可以将一个生成的新区块保存到chain.db

func (bc BlockChain) SaveData(data []byte) (Block, error) {
	db := bc.BoltDb
	//先查询chain.db中存储的最新的区块
	var e error
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			e = errors.New("桶未创建")
			return e
		}
		//lastHash := bucket.Get([]byte(LAST_KEY))
		lastBlockBytes := bucket.Get(bc.LastHash)
		lastBlock, _ = DeSerialize(lastBlockBytes)
		//if err != nil {
		//	panic(err.Error())
		//}
		return nil
	})
	//先生成一个区块，把data存入到新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1,data,lastBlock.Hash)

	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//key,value
		newBlockBytes, _ := newBlock.Serialze()
		bucket.Put(newBlock.Hash,newBlockBytes)
		//更新lasthash
		bucket.Put([]byte(LAST_KEY),newBlock.Hash)
		bc.LastHash = newBlock.Hash
		return nil
	})


	return newBlock, e
}
