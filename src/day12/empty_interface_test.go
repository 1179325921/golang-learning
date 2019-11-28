package day12

import (
	"fmt"
	"testing"
)

func TestEmptyInterface(t *testing.T) {
	slice := make([]interface{}, 10)
	m1 := make(map[string]string)
	m2 := make(map[int]int)
	m3 := make(map[string]map[string]string)
	m1["111"] = "111"
	m2[22] = 222
	m3["333"] = m1
	slice[0] = m1
	slice[1] = m2
	slice[3] = m3
	t.Log(slice)
}

func doSomething(op interface{}) {
	//if v, ok := op.(int); ok {
	//	fmt.Println("Integer", v)
	//	return
	//}
	//if v, ok := op.(string); ok {
	//	fmt.Println("string", v)
	//	return
	//}
	//fmt.Println("unknown")
	switch op.(type) {
	case int:
		fmt.Println("Integer", op)
	case string:
		fmt.Println("string", op)
	default:
		fmt.Println("unknown")
	}
}

func TestDoSomething(t *testing.T) {
	doSomething(1)
	doSomething(2.22)
}
