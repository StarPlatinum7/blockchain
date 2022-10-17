package main

import (
	"blockchain_example/core"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	//1.监听端口

	server, err := net.Listen("tcp", ":8087")

	if err != nil {
		fmt.Println("开启socket服务失败")
	}

	fmt.Println("正在开启 Server ...")

	for {
		//2.接收来自 client 的连接,会阻塞
		conn, err := server.Accept()

		if err != nil {
			fmt.Println("连接出错")
		}

		//并发模式 接收来自客户端的连接请求，一个连接 建立一个 conn，服务器资源有可能耗尽 BIO模式

		go connHandler(conn)
		fmt.Println("is")

	}

}
func connHandler(c net.Conn) {
	var get []string
	//1.conn是否有效
	if c == nil {
		log.Panic("无效的 socket 连接")
	}

	//2.新建网络数据流存储结构
	buf := make([]byte, 4096)
	//3.循环读取网络数据流
	for i := 0; i < 18; i++ {
		//3.1 网络数据流读入 buffer
		cnt, err := c.Read(buf)
		//3.2 数据读尽、读取错误 关闭 socket 连接
		if cnt == 0 || err != nil {
			c.Close()
			break
		}

		fCommand := string(buf[0:cnt])
		get = append(get, fCommand)

		fmt.Println("客户端传输->", fCommand)

		switch fCommand {

		default:
			c.Write([]byte("服务器端回复" + fCommand + "\n"))
		}

		fmt.Printf("来自 %v 的连接关闭\n", c.RemoteAddr())
	}
	err := c.Close()
	if err != nil {
		return
	} //关闭client端的连接，telnet 被强制关闭
	//get即为接收到的交易消息
	fmt.Println("接受交易信息成功")
	//初始化一个区块链实例
	myBlockchain := core.NewBlockchain(2)
	//向区块链中先后添加6笔交易，每笔交易间延时1~3秒，模拟交易间的先后顺序
	for i := 0; i < 18; i += 3 {
		value, _ := strconv.Atoi(get[i+2])
		myBlockchain.SendTransaction(get[i], get[i+1], value)
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	}
	//打印区块链中的交易历史
	myBlockchain.Print()
	//打印世界状态
	core.PrintAccount()

	return

}
