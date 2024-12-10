package main

import "fmt"

func noNameReturn() int {
	result := 10
	defer func() {
		result = 20
		fmt.Println("Defer called")
	}()
	return result // 返回值无法被 defer 修改
}

func namedReturn() (result int) {
	result = 10
	defer func() {
		result = 20 // 修改了命名返回值
	}()
	return result
}

func main() {
	fmt.Println(noNameReturn()) // 输出：10
	fmt.Println(namedReturn())  // 输出：20
}
