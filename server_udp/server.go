// server_tcp.go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 启动 UDP 服务端，监听端口 8080
	// 使用 :8080 来监听所有可用的网络接口
	address, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", address)
	if err != nil {
		fmt.Println("Error listening on UDP port:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("UDP Server listening on localhost:8080")

	buffer := make([]byte, 1024)
	for {
		// 从客户端接收数据
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// 打印接收到的原始字节数据
		fmt.Printf("Received data (raw bytes) from %s: %v\n", addr, buffer[:n])

		// 打印接收到的消息
		fmt.Printf("Received message: %s\n", string(buffer[:n]))

		// 向客户端发送响应
		message := "Message received"
		_, err = conn.WriteToUDP([]byte(message), addr)
		if err != nil {
			fmt.Println("Error sending to client:", err)
		}
	}
}
