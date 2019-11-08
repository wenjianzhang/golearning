package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	// 第一种赋值方式
	//var a int =1
	//var b int =1

	// 第二种赋值方式
	//var (
	//	a int = 1
	//	b     = 1
	//)

	// 第三种赋值方式
	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}

func TestExchang(t *testing.T) {
	a := 1
	b := 1
	// 1
	//tmp := a
	//a = b
	//b = tmp
	//
	// 2
	a, b = b, a
	t.Log(a, b)
}
