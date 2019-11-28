package main

import (
	"fmt"
	"sync"
)

func printDemo() {
	fmt.Println("demo函数")
}

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(printDemo)
			wg.Done()
		}()
	}
	wg.Wait()
}
