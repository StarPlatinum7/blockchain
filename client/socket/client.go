package socket

import (
	"fmt"
	"net"
	"strings"
)

func cConnHandler(c net.Conn) {

	//返回一个拥有 默认size 的reader，接收客户端输入
	//reader := bufio.NewReader(os.Stdin)
	//缓存 conn 中的数据
	buf := make([]byte, 1024)
	fmt.Println("请输入客户端请求数据...")
	sa := [][]string{
		{"Tom", "Mary", "20"},
		{"Tom", "Lily", "30"},
		{"Aron", "Tom", "10"},
		{"Byron", "Lily", "15"},
		{"Joe", "Tom", "20"},
		{"Lily", "Mary", "40"},
	}
	for _, temp := range sa {
		for _, post := range temp {
			//客户端输入
			//input, _ := reader.ReadString('\n')
			input := post
			//去除输入两端空格
			input = strings.TrimSpace(input)
			//客户端请求数据写入 conn，并传输
			c.Write([]byte(input))
			//服务器端返回的数据写入空buf
			cnt, err := c.Read(buf)

			if err != nil {
				fmt.Printf("客户端读取数据失败 %s\n", err)
				continue
			}

			//回显服务器端回传的信息
			fmt.Print("服务器端回复" + string(buf[0:cnt]))
		}

	}
}

func ClientSocket() {
	conn, err := net.Dial("tcp", "127.0.0.1:8087")
	if err != nil {
		fmt.Println("客户端建立连接失败")
		return
	}

	cConnHandler(conn)
}
