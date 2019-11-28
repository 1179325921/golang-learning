package main

import "fmt"

func main() {
	var a int = 8
	var b int = 100
	var ptr *int = &a
	*ptr = 10
	ptr = &b
	*ptr = 200
	fmt.Printf("a=%d, b=%d, *ptr=%d", a, b, *ptr)
}
