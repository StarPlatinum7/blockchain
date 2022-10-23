package core

import (
	"encoding/hex"
	"math/rand"
	"time"
)

type accounts struct {
	account []*account
}
type account struct {
	name    string
	address string
	balance int
}

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

// 请添加结构存储账户，要求其中至少包含两个字段：账户的地址和余额（Balance）。给交易中出现的每个人名分配一个账户，地址可以自由指定，
// 余额设置成100。改写对应的函数和方法，要注意在交易**最终被提交到区块链上**时，发出方和接受方的余额也要进行相应变动。创建一个列表，
// 存储系统中所有账户地址，然后，添加一个输出世界状态函数，输出世界状态。“世界状态”就是经过目前已提交到区块链上的所有交易后，各账户的余额
func (person *account) more(money int) {
	person.balance -= money
}
func (person *account) less(money int) {
	person.balance -= money
}

// 先根据交易情况初始化一个账户

var accountts accounts

func GenerateNewAccount(name string) *account {
	temp := new(account)
	temp.name = name
	temp.balance = 100
	temp.address = RandStringBytesMaskImprSrc(8) //生成8位数的地址
	return temp
}
func (accounts *accounts) AppendAccount(NewAccount *account) {
	isappend := false
	if len(accounts.account) == 0 {
		accounts.account = append(accounts.account, NewAccount)
	} else {
		for _, temp := range accounts.account {
			if temp.name == NewAccount.name {
				isappend = true //已经加过
			}
		}
		if !isappend {
			accounts.account = append(accounts.account, NewAccount)
		}
	}
}

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, (n+1)/2) // can be simplified to n/2 if n is always even

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)[:n]
}
