package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex

func WriteData(wg *sync.WaitGroup, n int) {
	rw.Lock()
	fmt.Printf("增加第%d个写锁\n", n)
	for i := 0; i < 10; i++ {
		fmt.Println("正在进行写操纵")
		//写携程休眠
		time.Sleep(time.Millisecond * 20)
	}
	fmt.Printf("释放第%d个写锁\n", n)
	rw.Unlock()
	wg.Done()
}
func ReadData(wg *sync.WaitGroup, n int) {
	rw.RLock()
	fmt.Sprintf("增加第%d个读锁\n", n)
	for i := 0; i < 10; i++ {
		fmt.Println("正在进行读操作")
		//读携程休眠
		time.Sleep(time.Millisecond * 20)
	}
	fmt.Printf("释放第%d个读锁\n", n)
	rw.RUnlock()
	wg.Done()
}

func main() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go WriteData(&wg, i)
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go ReadData(&wg, i)
	}
	wg.Wait()
}
