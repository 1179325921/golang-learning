package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("参数共有:%d\n", len(os.Args))
	for key, value := range os.Args {
		fmt.Printf("args[%d],%v\n", key, value)
	}
}
