package day06

import "testing"

func TestSliceInit(t *testing.T) {
	var slice1 []int
	t.Log(len(slice1), cap(slice1))
	slice1 = append(slice1, 1)
	t.Log(len(slice1), cap(slice1))

	slice2 := []int{1, 2, 3, 4}
	t.Log(len(slice2), cap(slice2))

	slice3 := make([]int, 3, 5)
	t.Log(len(slice3), cap(slice3))
	t.Log(slice3[0], slice3[1], slice3[2])
	slice3 = append(slice3, 1)
	t.Log(slice3[0], slice3[1], slice3[2], slice3[3])
	t.Log(len(slice3), cap(slice3))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

//func TestSliceCompare(t *testing.T){
//	s1:=[]int{1,2}
//	s2:=[]int{1,2}
//	if s1==nil{
//
//	}
//}
