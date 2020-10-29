package blockchain

import (
	"DataCertProject/util"
	"bytes"
	"crypto/sha256"
	"math/big"
)

//右移几位，对应难度
const DIFFFICULTY  = 16
/*
	工作量证明结构体

*/
type ProofOfWork struct {
	//目标值
	Target *big.Int
	//区块信息
	Block Block
}

/**
	实例化一个pow算法实例
	区块信息	 未知值		 目标值
	  A   +   nonce   <    b

*/

//求出【目标值】
func NewPoW(block Block) ProofOfWork {
	//
	target := big.NewInt(1)
	target.Lsh(target, 255-DIFFFICULTY)
	pow :=  ProofOfWork{
		Target: target,
		Block: block,
	}
	return pow
}

//pow算法：寻找nonce值
/**
	要寻找nonce值，先将【区块信息】的值求出来

	Height int64		//区块高度
	TimeStamp int64		//时间戳
	Hash []byte			//区块的hash
	Data []byte			//区块数据
	PrevHash []byte		//上一个区块的哈希
	Version string		//版本号
	Nonce int64			//随机数
*/

func (p ProofOfWork) Run() ([]byte,int64) {
	var nonce int64
	//var bigBlock *big.Int//用于接收由[]byte转为*big.Int
	bigBlock := new(big.Int)
	var block256Hash []byte
	for {
		block := p.Block

		heightBytes, _ :=util.IntToBytes(block.Height)
		timeBytes, _ := util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		nonceBytes, _ := util.IntToBytes(nonce)//nonce不断变化

		//拼接【区块信息】
		blocBytes := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
			//暂时无法操作hash，因为还正在找nonce，hash是一个区块的全部hash，包括nonce
		},[]byte{})
		sha256Hash := sha256.New()
		sha256Hash.Write(blocBytes)
		block256Hash = sha256Hash.Sum(nil)

		//fmt.Printf("挖矿中，当前尝试nonce值:%d\n",nonce)
		//【区块信息】 + nonce 对应的*big.Int
		bigBlock = bigBlock.SetBytes(block256Hash)
		//fmt.Printf("目标值：%x\n,hash值：%x\n",p.Target,bigBlock)
		if p.Target.Cmp(bigBlock) == 1 {//1 > , 0 = ,-1 <
			//找到合适的nonce,退出循环
			break
		}
		nonce++//条件不满足，nonce+1，继续循环
		//time.Sleep(100)
	}
	return block256Hash,nonce
}