// server_tcp.go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 启动 TCP 服务端，监听端口 8080
	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listen.Close()

	fmt.Println("Server listening on localhost:8081")

	// 接受客户端连接，并持续处理每个连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 在新的 goroutine 中处理每个连接
		go handleRequest(conn)
	}
}

// 处理每个连接的请求
func handleRequest(conn net.Conn) {
	// 关闭连接
	defer conn.Close()

	// 循环读取客户端发送的消息，直到客户端关闭连接
	for {
		// 读取客户端发送的数据
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from client_tcp:", err)
			return
		}

		// 打印接收到的原始字节数据
		fmt.Printf("Received data (raw bytes): %v\n", buffer[:n])

		// 打印接收到的消息
		fmt.Printf("Received message: %s\n", string(buffer[:n]))

		// 向客户端发送响应
		message := "Message received"
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing to client_tcp:", err)
			return
		}
	}
}
