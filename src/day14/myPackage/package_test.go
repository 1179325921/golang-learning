package myPackage

import (
	s "day14/series"
	"fmt"
	"testing"
)

func init() {
	fmt.Println("myInit3")
}

func TestPackage(t *testing.T) {
	t.Log(s.GetFibonacci(2))
	//小写的方法名无法调用
	//s.get(1)
}
