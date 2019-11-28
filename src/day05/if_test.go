package day05

import "testing"

func TestIf(t *testing.T) {
	if n := 1 == 1; n {
		t.Log("1==1")
	}
}

func TestIfMultiSec(t *testing.T) {
	if result, err := someFun(); err == nil {
		t.Log("出错啦")
	} else {
		t.Log(result)
	}
}
