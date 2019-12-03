package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(srcFile, destFile string) error {
	file1, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("打开文件失败")
		return err
	}
	reader := bufio.NewReader(file1)
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("创建文件失败")
		return err
	}
	writer := bufio.NewWriter(file2)
	_, err = io.Copy(writer, reader)
	if err != nil {
		return err
	}
	defer file1.Close()
	defer file2.Close()
	return nil
}

func main() {
	err := CopyFile("a.txt", "c.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("copy文件成功")
	}
}
