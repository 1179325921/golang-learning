package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, e := ioutil.ReadFile("/Users/admin/workspace/golang-learning/src/file/demo1/a.txt")
	if nil != e {
		fmt.Println(e)
	}
	fmt.Println(string(bytes))
}
