package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
	读取文件内容，并写入到新文件中
*/
func main() {
	//　读取文件
	bytes, e := ioutil.ReadFile("a.txt")
	if e != nil {
		fmt.Println("读取文件失败", e)
		return
	}
	// 写入文件
	b, e := PathExists("b.txt")
	if b {
		err := ioutil.WriteFile("b.txt", bytes, 0666)
		if err != nil {
			fmt.Println("写入文件失败", err)
			return
		}
	} else {
		if e == nil {
			// 创建文件
			file, err := os.OpenFile("b.txt", os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				fmt.Println("创建文件失败", err)
				return
			}
			file.Write(bytes)

		} else {
			fmt.Println(e)
			return
		}
	}

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
