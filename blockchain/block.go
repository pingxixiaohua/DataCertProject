package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

//区块结构体的定义
type Block struct {
	Height int64		//区块高度
	//Size int			//区块大小，算出来的
	TimeStamp int64		//时间戳
	Hash []byte			//区块的hash
	Data []byte			//区块数据
	PrevHash []byte		//上一个区块的哈希
	Version string		//版本号
	Nonce int64			//随机数，用于pow工作量证明算法计算
}

func CreateGenesisBlock() Block {
	block := NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	return block
}

//新建一个区块，
func NewBlock(height int64,data []byte, prevHash []byte) (Block) {
	//1、构建一个block实例，用于生成区块
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Version:   "0x01",
	}

	//2、为新生成的区块找nonce
	pow := NewPoW(block)
	blockHash,nonce := pow.Run()

	//3、将block的Nonce设置为找到的合适的nonce数
	block.Nonce = nonce


	/*	util.SHA256Hash要求一个[]byte参数
			block中包含3个已是[]byte
			将Height、TimeStam、Version转[]byte格式
			最后将6个[]byte进行拼接

	*/

	//调用SHA256Hash进行hash计算
	//4、设置第七个字段Hash
	//block.Hash = util.SHA256Hash(blockBytes)
	block.Hash = blockHash

	return block
}

/*
	区块的序列化
*/
func (bk Block) Serialze() ([]byte, error){
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(bk)
	if err != nil {
		return nil,err
	}
	return buff.Bytes(),nil
}

/*
	区块的反序列化
*/

func DeSerialize(data []byte) (*Block, error) {
	var block Block
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err != nil {
		return nil,err
	}
	return &block,nil
}