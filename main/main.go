package main

import (
	"blockchain_example/core"
)

func main() {
	//初始化一个区块链实例
	myBlockchain := core.NewBlockchain(2)
	//向区块链中先后添加6笔交易，每笔交易间延时1~3秒，模拟交易间的先后顺序
	myBlockchain.SendTransaction("Tom", "Mary", 20)
	//time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Tom", "Lily", 30)
	//time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Aron", "Tom", 10)
	//time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Byron", "Lily", 15)
	//time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Joe", "Tom", 20)
	//time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Lily", "Mary", 40)

	//打印区块链中的交易历史
	myBlockchain.Print()
	//打印世界状态
	core.PrintAccount()
}
