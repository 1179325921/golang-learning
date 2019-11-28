package day09

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//函数多返回值
func returnMultiValues() (int, int) {
	return rand.Intn(5), rand.Intn(20)
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

//inner函数作为参数，返回值也是一个函数
func timeSpent(inner func(op int) int) func(op int) int {
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

func varArgs(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarArgs(t *testing.T) {
	t.Log(varArgs(1, 2, 3, 4))
	t.Log(varArgs(1, 2, 3, 4, 100))
}

func Clear() {
	fmt.Println("释放资源")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("哈哈哈")
	panic("err")
}
