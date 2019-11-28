package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("hello world!!!")
}

func main() {
	//fmt.Println("main execution started")
	//// create goroutines
	//go printHello()
	//
	////sleep
	//time.Sleep(time.Second * 20)
	//fmt.Println("main execution stopped")
	for i := 0; i < 10; i++ {
		//匿名携程
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	//main携程休眠
	time.Sleep(time.Millisecond * 20)
}
