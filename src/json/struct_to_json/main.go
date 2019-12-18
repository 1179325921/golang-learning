package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Phone string
}

func main() {
	s := Student{
		Name:  "yuanshuai",
		Age:   25,
		Phone: "111111111111",
	}
	marshal, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("s序列化失败:%v", marshal)
		return
	}
	fmt.Println(string(marshal))
}
