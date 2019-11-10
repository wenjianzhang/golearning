package err_test

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")

var LargerhanHundredError = errors.New("n should be not less than 2")

func GetFibonaccin(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerhanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonaccin(t *testing.T) {
	if v, err := GetFibonaccin(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less")
		}
		if err == LargerhanHundredError {

		}
		t.Log(err)
	} else {
		t.Log(v)
	}

}
