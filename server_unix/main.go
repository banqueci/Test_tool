package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 定义 Unix Socket 的地址
	socketPath := "/tmp/my_unix_socket"

	// 尝试删除现有的 Unix Socket 文件
	_ = os.Remove(socketPath)

	// 创建一个 Unix 套接字地址
	addr := &net.UnixAddr{Name: socketPath, Net: "unix"}

	// 创建并监听 Unix 套接字
	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		fmt.Println("Error listening on Unix socket:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on", socketPath)

	// 持续接收客户端消息
	for {
		// 等待并接受客户端连接
		conn, err := listener.AcceptUnix()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 处理客户端连接
		go handleConnection(conn)
	}
}

func handleConnection(conn *net.UnixConn) {
	defer conn.Close()

	// 创建一个缓冲区接收数据
	buf := make([]byte, 1024)

	for {
		// 读取客户端发送的消息
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		// 打印接收到的消息
		message := string(buf[:n])
		fmt.Printf("Received: %s\n", message)

		// 根据接收到的消息返回响应
		var response string
		if message == "exit" {
			response = "Goodbye!"
		} else {
			response = "Message received: " + message
		}

		// 发送响应到客户端
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}
