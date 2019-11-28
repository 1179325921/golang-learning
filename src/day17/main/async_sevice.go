package main

import (
	"fmt"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("do something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is Done")
}

func asyncService() chan string {
	////串行
	//fmt.Println(service())
	//otherTask()
	//retCh := make(chan string)
	//定义一个buffered channel
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func main() {
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)
}
