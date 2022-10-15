package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

// Blockchain 是一个简易的"区块链"结构体，区块链由多个区块构成
//
// # Blocks 是由区块组成的切片，代表着区块链中已经上链的区块
//
// PendingTransactions 是用户已经提交但是未上链的交易组成的切片，之所以未上链，是因为交易数量还不够，只有当交易数量达到每个区块最大可容纳交易数量
// MaxTransactionAmountPerBlock 时，这些交易才会被一并打包成一个区块并上链
type Blockchain struct {
	Blocks                       []*Block
	PendingTransactions          []*transaction
	MaxTransactionAmountPerBlock int
}

// 用来存放一次交易的信息
type transaction struct {
	From            string
	To              string
	Value           int
	FullName        string
	TransactionHash string
}

// NewBlockchain 会初始化一个 Blockchain 结构体并返回，该区块链每个区块的最大交易数量会被设置为传入的参数maxTransactionAmountPerBlock
//
// 函数执行的流程为：
//
// (1)生成一个 Blockchain 区块链实例，并将该实例的字段初始化
//
// (2)使用函数 GenerateGenesisBlock 生成一个创世区块
//
// (3)使用方法 AppendNewBlock 将创世区块添加到该区块链实例中
//
// (4)返回该区块链实例
func NewBlockchain(maxTransactionAmountPerBlock int) *Blockchain {
	blockchain := new(Blockchain)
	blockchain.MaxTransactionAmountPerBlock = maxTransactionAmountPerBlock
	genesisBlock := GenerateGenesisBlock()
	blockchain.AppendNewBlock(genesisBlock)
	return blockchain
}

// SendTransaction 会处理每个传入进来的交易，具体的处理流程为：
//
// (1)将交易添加到区块链的待上链交易中
//
// (2)判断待处理交易数量，如果数量大于前面设置的每个区块最大可容纳交易数量，则使用 GenerateNewBlock 函数生成一个新区块，再使用 AppendNewBlock 函数将该区块添加到区块链中，否则不作任何操作
func (bc *Blockchain) SendTransaction(from string, to string, value int) {

	//先初始化这次交易
	NewTran := new(transaction)
	NewTran.From = from
	NewTran.To = to
	NewTran.Value = value
	NewTran.FullName = from + " gave " + to + " " + string(value) + " dollar"
	//对应此次交易的hash
	record := NewTran.FullName + string(time.Now().Unix())
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	hashValue := hex.EncodeToString(hashed)
	//返回生成的Hash
	NewTran.TransactionHash = hashValue //存储

	bc.PendingTransactions = append(bc.PendingTransactions, NewTran)
	if len(bc.PendingTransactions) >= bc.MaxTransactionAmountPerBlock {
		bc.AppendNewBlock(GenerateNewBlock(bc.Blocks[len(bc.Blocks)-1], bc.PendingTransactions))
		bc.PendingTransactions = []*transaction{}
	}
}

// AppendNewBlock 会将传入的新区块添加到区块链中，具体的处理流程为：
//
// (1)如果区块链中目前没有区块，则直接将新区块（也即创世区块）添加到区块链中
//
// (2)如果区块链中目前已经有区块，则需要将新区块传入 CheckValid 函数，判断新区块是否合法（合法的定义可见 CheckValid 函数注释），若合法，则将新区块添加到区块链中，否则向终端报错并终止程序运行
func (bc *Blockchain) AppendNewBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		if bc.CheckValid(newBlock) {
			bc.Blocks = append(bc.Blocks, newBlock)
		} else {
			log.Fatal("invalid new block")
		}
	}
}

// CheckValid 返回新区块的合法性
//
// 新区块只有同时满足以下条件时才是合法的：
//
// (1)新区块的序号 Index 等于前一区块的序号值加一
//
// (2)新区块的 PreviousHash 字段值等于前一区块的 Hash 字段值
//
// (3)新区块使用 CalculateHash 函数计算出来的哈希结果等于新区块的 Hash 字段值
func (bc *Blockchain) CheckValid(newBlock *Block) bool {
	oldBlock := bc.Blocks[len(bc.Blocks)-1]
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PreviousHash != oldBlock.Hash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// Print 可以打印出来区块链的所有历史
func (bc *Blockchain) Print() {
	for i := range bc.Blocks {
		fmt.Println("Block Index", bc.Blocks[i].Index)
		fmt.Println("\t", "-timestamp: ", bc.Blocks[i].TimeStamp)
		fmt.Println("\t", "-previous hash: ", bc.Blocks[i].PreviousHash)
		fmt.Println("\t", "-hash: ", bc.Blocks[i].Hash)
		fmt.Println("\t", "-transactions:")
		for j := range bc.Blocks[i].Transactions {
			fmt.Println("\t\t", "--transaction", j, "of block", i, ":", bc.Blocks[i].Transactions[j].FullName)
		}
	}
}
