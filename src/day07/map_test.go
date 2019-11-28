package day07

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(len(m1))
	t.Log(m1["two"])
	m2 := map[int]int{}
	m2[4] = 16
	t.Log(len(m2))
	t.Log(m2[4])
	m3 := make(map[int]int, 10)
	t.Log(m3)
}

func TestAccessNotExitKey(t *testing.T) {
	m1 := make(map[int]int, 10)
	t.Log(m1[0], m1[2])
	m1[2] = 0
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Log("key3存在", v)
	} else {
		t.Log("key3不存在", v)
	}
}

func TestRangeMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
func TestMapWithFuncValue(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return op * op * op }
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}

func TestMapForSet(t *testing.T) {
	m1 := map[int]bool{}
	m1[1] = true
	n := 3
	if m1[n] {
		t.Logf("%d is exit", n)
	} else {
		t.Logf("%d is not exit", n)
	}
	t.Log(len(m1))
	delete(m1, 1)
	n = 1
	if m1[n] {
		t.Logf("%d is exit", n)
	} else {
		t.Logf("%d is not exit", n)
	}
}
