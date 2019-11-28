package day13

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoErr = errors.New("n should not less than 2")

var LargeThan100Err = errors.New("n should not less than 2")

//斐波那契数列
func getFibonacci(n int) ([]int, error) {
	if n < 2 {
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

func TestGetFibonacci(t *testing.T) {
	if v, err := getFibonacci(-2); err != nil {
		if err == LessThanTwoErr {
			fmt.Println("less than")
		}
		if err == LargeThan100Err {
			fmt.Println("too large")
		}
	} else {
		t.Log(v)
	}
}
