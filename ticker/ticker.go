package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second) // 每秒触发一次
	defer ticker.Stop()                       // 确保程序退出时释放定时器资源

	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second) // 模拟 5 秒后停止程序
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-done:
			fmt.Println("Ticker stopped")
			return
		}
	}
}
