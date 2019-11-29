package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, e := os.OpenFile("a.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if nil != e {
		fmt.Println(e)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "你好，golang\r\n"
	for i := 0; i < 5; i++ {
		_, e := writer.WriteString(str)
		if e != nil {
			fmt.Println(e)
			return
		}
	}
	e = writer.Flush()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("结束")
}
