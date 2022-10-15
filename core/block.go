package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block 是一个简易的"区块"结构体，区块由多笔交易构成
//
// # Index 是该区块的序号，在区块链中，区块的序号单调递增，创世区块的序号为0
//
// # TimeStamp 是该区块的时间戳
//
// # PreviousHash 是前一区块的哈希值，正是因为这一字段的存在，各个区块才能组成防篡改的链式结构
//
// # Hash 是当前区块的哈希值，这一字段保证了本区块内的交易不被篡改
//
// Transactions 是当前区块内的交易组成的切片，是该结构体的主要荷载
type Block struct {
	Index        int64
	TimeStamp    int64
	PreviousHash string
	Hash         string
	Transactions []string
}

// GenerateNewBlock 会生成一个新的区块并返回，具体执行流程包括：
//
// (1)生成一个 Block 区块实例，并将该实例的字段初始化
//
// (2)该实例中的 Index 字段和 PreviousHash 的初始化需要借助传参previousBlock的信息
//
// (2)该实例中的 Transactions 字段由传参 transactions 得到
//
// (4)通过 CalculateHash 函数计算该区块的哈希值，并将结果赋给该实例的 Hash 字段
//
// (5)返回该区块实例
func GenerateNewBlock(previousBlock *Block, transactions []string) *Block {
	block := new(Block) //定义一个新块

	block.Index = previousBlock.Index + 1
	block.TimeStamp = time.Now().Unix()
	block.PreviousHash = previousBlock.Hash
	block.Hash = CalculateHash(block)
	block.Transactions = transactions
	return block
}

// CalculateHash 计算传入的结构体的哈希值，以字符串形式返回
func CalculateHash(b *Block) string {
	var hashValue string

	record := string(b.Index) + string(b.TimeStamp) + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	//返回生成的Hash
	hashValue = hex.EncodeToString(hashed)

	return hashValue
}

// GenerateGenesisBlock 的作用与 GenerateNewBlock 类似，都是生成新区块，但是本函数用于生成特殊的区块“创世区块”，创世区块是区块链中的第一个区块，
// 它的序号（Index字段值）为0
func GenerateGenesisBlock() *Block {
	block := new(Block)
	block.Index = 0
	block.PreviousHash = ""
	block.TimeStamp = time.Now().Unix()
	block.Transactions = []string{"***genesis block***"}
	block.Hash = CalculateHash(block)
	return block
}
