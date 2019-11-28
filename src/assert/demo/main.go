package main

import "fmt"

func main() {
	// 类型断言
	var x interface{}
	b := 2.1
	x = b
	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是 %T,值=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行")
}
