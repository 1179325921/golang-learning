package day11

import (
	"fmt"
	"testing"
	"time"
)

//定义函数别名
type myIntFun func(op int) int

//inner函数作为参数，返回值也是一个函数
func timeSpent(inner myIntFun) myIntFun {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func timeSlow(n int) int {
	time.Sleep(time.Second * 2)
	return n
}

func TestTimeSpent(t *testing.T) {
	ref := timeSpent(timeSlow)(10)
	t.Log(ref)
}
