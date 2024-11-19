// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 解析服务器的地址
	address, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	// 连接到 UDP 服务端
	conn, err := net.DialUDP("udp", nil, address)
	if err != nil {
		fmt.Println("Error connecting to server_tcp:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 持续接收用户输入，直到用户输入 "exit" 退出
	for {
		// 读取用户输入
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message to send to server_tcp (or type 'exit' to quit): ")
		message, _ := reader.ReadString('\n')

		// 如果输入 "exit" 则退出
		if message == "exit\n" {
			fmt.Println("Exiting client...")
			break
		}

		// 向服务端发送数据
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to server_tcp:", err)
			return
		}

		// 接收服务端的响应
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server_tcp:", err)
			return
		}

		// 打印接收到的响应
		fmt.Printf("Server response: %s\n", string(buffer[:n]))
	}
}
