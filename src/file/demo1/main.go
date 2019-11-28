package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, e := os.Open("/Users/admin/workspace/golang-learning/src/file/demo1/a.txt")
	if nil != e {
		fmt.Println(e)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		s, e := reader.ReadString('\n')
		if e == io.EOF {
			break
		}
		fmt.Print(s)
	}
	fmt.Println("读取文件结束")
}
