package main

import (
	"fmt"
	"time"
)

//1个发送者1个监听者(发送端直接关闭)

func worker(ch chan int) {
	go func() {
		for {
			if v, ok := <-ch; ok {
				fmt.Println(v)
			} else {
				fmt.Println("接收不到数据了！！")
				break
			}
		}
		time.Sleep(time.Millisecond * 50)
	}()

}

func sender(ch chan int) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
		time.Sleep(time.Millisecond * 50)
	}()
}

func main() {
	ch := make(chan int, 100)
	sender(ch)
	time.Sleep(time.Millisecond * 50)
	worker(ch)
	time.Sleep(time.Millisecond * 50)
}
