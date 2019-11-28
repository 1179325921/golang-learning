package test

import "testing"

type MyInt int32

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64 = 2
	a = int32(b)
	var c MyInt
	c = MyInt(a)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr+1
	t.Log(aPtr)
	t.Logf("%T,%T", a, aPtr)
}

func TestStr(t *testing.T) {
	var str string
	t.Log("*" + str + "*")
	t.Log(len(str))
	if str == "" {
		t.Log("str是空字符串")
	}
}
