package series

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("myInit1")
}

func init() {
	fmt.Println("myInit2")
}

var LessThanTwoErr = errors.New("n should not less than 2")

var LargeThan100Err = errors.New("n should not less than 2")

//斐波那契数列
func GetFibonacci(n int) ([]int, error) {
	if n <= 2 {
		return nil, LessThanTwoErr
	}
	if n > 100 {
		return nil, LargeThan100Err
	}
	firstList := []int{1, 1}
	for i := 2; i <= n; i++ {
		firstList = append(firstList, firstList[n-2]+firstList[n-1])
	}
	return firstList, nil
}

func get(n int) {
	fmt.Println(n)
}
