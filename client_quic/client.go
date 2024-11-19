package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go"
	"log"
	"os"
)

func main() {
	// 创建上下文
	ctx := context.Background()

	// QUIC 配置（可选）
	quicConfig := &quic.Config{}

	// 建立一个 QUIC 会话
	session, err := quic.DialAddr(ctx, "localhost:4242", &tls.Config{InsecureSkipVerify: true}, quicConfig)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}
	fmt.Println("Connected to server. Type your messages below:")

	// 打开一个流
	stream, err := session.OpenStreamSync(ctx)
	if err != nil {
		log.Fatalf("Failed to open stream: %v\n", err)
	}

	// 使用 bufio.Writer 包装流
	writer := bufio.NewWriter(stream)

	// 使用 bufio.Scanner 读取用户输入
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // 读取失败（如 EOF）时退出
		}
		message := scanner.Text()

		// 检查退出命令
		if message == "exit" {
			fmt.Println("Exiting client.")
			break
		}

		// 写入消息到流
		_, err := writer.WriteString(message) // 添加换行符方便服务器解析
		if err != nil {
			log.Printf("Failed to write message: %v\n", err)
			break
		}

		// 刷新缓冲区，确保数据被发送
		err = writer.Flush()
		if err != nil {
			log.Printf("Failed to flush buffer: %v\n", err)
			break
		}
		fmt.Printf("Sent: %s\n", message)
	}

	// 关闭流和会话
	err = stream.Close()
	if err != nil {
		log.Printf("Failed to close stream: %v\n", err)
	}
	err = session.CloseWithError(0, "Client closed")
	if err != nil {
		log.Printf("Failed to close session: %v\n", err)
	}
}
