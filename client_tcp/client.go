// client_tcp.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接到 TCP 服务端
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 持续接受用户输入直到用户选择退出
	for {
		// 读取用户输入
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message to send to server_tcp (or type 'exit' to quit): ")
		message, _ := reader.ReadString('\n')

		// 如果用户输入 "exit" 则退出循环
		if message == "exit\n" {
			fmt.Println("Exiting client_tcp...")
			break
		}

		// 打印发送的原始字节数据
		fmt.Printf("Sending data (raw bytes): %v\n", []byte(message))

		// 向服务端发送用户输入的消息
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to server_tcp:", err)
			return
		}

		// 接收服务端的响应
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from server_tcp:", err)
			return
		}

		// 打印接收到的原始字节数据
		fmt.Printf("Received data (raw bytes): %v\n", buffer)

		// 打印服务端的响应
		fmt.Println("Server response:", string(buffer))
	}
}
