package main

import (
	"blockchain_example/core"
	"math/rand"
	"time"
)

func main(){
	//初始化一个区块链实例
	myBlockchain:=core.NewBlockchain(2)

	//向区块链中先后添加6笔交易，每笔交易间延时1~3秒，模拟交易间的先后顺序
	myBlockchain.SendTransaction("Tom gave Mary 20 dollars.")
	time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Tom gave Lily 30 dollars.")
	time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Aron gave Tom 10 dollars.")
	time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Byron gave Lily 15 dollars.")
	time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Joe gave Tom 20 dollars.")
	time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
	myBlockchain.SendTransaction("Lily gave Mary 40 dollars.")

	//打印区块链中的交易历史
	myBlockchain.Print()
}
