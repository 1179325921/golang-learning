package day05

import "testing"

func TestSwitch(t *testing.T) {
	for n := 0; n < 5; n++ {
		switch n {
		case 0, 1:
			t.Log("0-1")
		case 2, 3:
			t.Log("2-3")
		default:
			t.Log("不属于0-3")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for n := 0; n < 5; n++ {
		switch {
		case n%2 == 0:
			t.Log("偶数")
		case n%2 == 1:
			t.Log("奇数")
		default:
			t.Log("不知道什么数")
		}
	}

}
