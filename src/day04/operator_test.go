package day04

//运算符 比较数组
import "testing"

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{2, 4, 6, 8}
	c := [...]int{1, 2, 3, 4}
	//d:=[...]int{1,1,1,1,1,1}
	t.Log(a == b)
	t.Log(a == c)
}
