package test

import "testing"

func TestFibonacci(t *testing.T) {
	//定义变量
	//var a int = 1
	//var b int = 1
	//var (
	//	a int = 1
	//	b     = 1
	//)
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	//tmp:=a
	//a=b
	//b=tmp
	//一个语句可赋值多个变量
	a, b = b, a
	t.Log(a, b)
}
