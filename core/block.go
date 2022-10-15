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

// MerkleNode MerkleTree 节点
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// MerkleTree 根
type MerkleTree struct {
	RootNode *MerkleNode
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

	//data := make([][]byte, len(b.Transactions))
	//for i := range data {
	//	data[i] = make([]byte, len([]byte(b.Transactions[0])))
	//}
	//for _, temp := range b.Transactions {
	//	data = append(data, []byte(temp))
	//}

	//正确的代码应该是上面括号内的，但是总有错

	data := [][]byte{[]byte("b.Transactions[0]"), []byte("b.Transactions[1]")}

	tree := NewMerkleTree(data)
	hashValue = hex.EncodeToString(tree.RootNode.Data)
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

// NewMerkleNode 创建Merkle节点
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}
	// 创建存储明文信息的叶子节点
	if left == nil && right == nil {
		node.Data = data
		// 创建只有一个分支的MerkleNode
	} else if left != nil && right == nil {
		hash := sha256.Sum256(left.Data)
		node.Data = hash[:]
		// 创建有两个分支的MerkleNode
	} else {
		// slice = append(slice, anotherSlice...) 两个slice拼接在一起时要加...
		hash := sha256.Sum256(append(left.Data, right.Data...))
		node.Data = hash[:]
	}
	node.Left = left
	node.Right = right

	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	// 将所有数据构建为dataNode节点，接入node节点的左分支，并将node节存到nodes数组中
	for _, datum := range data {
		dataNode := NewMerkleNode(nil, nil, datum)
		node := NewMerkleNode(dataNode, nil, nil)
		nodes = append(nodes, *node)
	}

	for {
		var newLevel []MerkleNode

		// 根据当前层的节点，构造上一层
		// 当前层节点为奇数时
		if len(nodes)%2 == 1 {
			for j := 0; j < len(nodes)-1; j += 2 {
				node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
				newLevel = append(newLevel, *node)
			}
			node := NewMerkleNode(&nodes[len(nodes)-1], nil, nil)
			newLevel = append(newLevel, *node)
			// 当前层节点为偶数时
		} else {
			for j := 0; j < len(nodes); j += 2 {
				node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
				newLevel = append(newLevel, *node)
			}
		}

		// 更新层节点
		nodes = newLevel
		if len(nodes) == 1 {
			break
		}
	}
	mTree := MerkleTree{&nodes[0]}
	return &mTree
}
