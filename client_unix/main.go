package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 定义 Unix Socket 的地址（与服务器端一致）
	socketPath := "/tmp/my_unix_socket"

	// 连接到 Unix 套接字
	addr := &net.UnixAddr{Name: socketPath, Net: "unix"}
	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		fmt.Println("Error connecting to Unix socket:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 创建一个输入扫描器，读取命令行输入
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your message to send (press Enter to send, 'exit' to quit):")

	// 无限循环，等待用户输入
	for {
		// 读取用户输入
		scanner.Scan()
		message := scanner.Text()

		// 如果用户输入 'exit'，退出程序
		if message == "exit" {
			fmt.Println("Exiting client.")
			break
		}

		// 发送消息到服务器
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
		fmt.Printf("Sent: %s\n", message)

		// 接收服务器的响应
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		// 打印服务器的响应
		fmt.Printf("Received: %s\n", string(buf[:n]))
	}
}
