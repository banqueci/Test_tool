package main

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"github.com/quic-go/quic-go"
)

func main() {
	// 打开密钥日志文件
	keyLogFile, err := os.OpenFile("keylogfile.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("无法打开密钥日志文件: %v\n", err)
	}
	defer keyLogFile.Close()

	// 配置 TLS，启用密钥记录
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,       // 跳过服务器验证（测试用）
		KeyLogWriter:       keyLogFile, // 记录 TLS 密钥
	}

	// 配置 QUIC
	quicConfig := &quic.Config{}

	// 创建 QUIC 连接
	ctx := context.Background()
	session, err := quic.DialAddr(ctx, "localhost:4242", tlsConfig, quicConfig)
	if err != nil {
		log.Fatalf("无法连接到服务器: %v\n", err)
	}
	_ = session
	
	log.Println("QUIC 会话建立成功，密钥已记录到 keylogfile.log")
}
