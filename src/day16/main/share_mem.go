package main

import (
	"fmt"
	"sync"
)

//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 6; i++ {
//		//计数器加一
//		wg.Add(1)
//		go func(i int) {
//			fmt.Printf("第%d个携程执行结束\n", i)
//			//计数器减一
//			wg.Done()
//		}(i)
//	}
//	//计数器不为零时，一直阻塞
//	wg.Wait()
//	fmt.Println("main goroutine执行结束")
//}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(2)
	sum := 0
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock() //上锁
			sum++
			mutex.Unlock() //解锁
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()
			sum++
			mutex.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(sum)
}
