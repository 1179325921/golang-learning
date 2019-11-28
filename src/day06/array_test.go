package day06

import "testing"

func TestArrayInit(t *testing.T) {
	arr1 := [3]int{}
	arr2 := [2]string{"你", "好"}
	t.Log(arr1)
	t.Log(arr2)
	arr3 := [...]int{2, 4, 6, 8, 10}
	t.Log(arr3, len(arr3))
}

func TestArrayTraversal(t *testing.T) {
	arr3 := [...]int{2, 4, 6, 8, 10}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	//range 取值和索引
	for index, value := range arr3 {
		t.Log(index, value)
	}
	//只取值
	for _, value := range arr3 {
		t.Log(value)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 3, 5, 7, 9}
	t.Log(arr3[3:])
	t.Log(arr3[:])
	t.Log(arr3[:4])
}
