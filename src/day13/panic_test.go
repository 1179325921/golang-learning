package day13

import (
	"errors"
	"fmt"
	"testing"
)

func TestFanic(t *testing.T) {
	//出现错误也会执行
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("recover from ", error)
		}
	}()
	fmt.Println("start")
	panic(errors.New("something error"))
}

//func TestExit(t *testing.T) {
//	// 不会执行
//	defer func() {
//		fmt.Println("finally")
//	}()
//	fmt.Println("start")
//	os.Exit(-1)
//}
