package mianshi

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Add(a int, b int) int {
	c := a + b
	fmt.Println(a, b, c)
	return c
}

func TestAdd(t *testing.T) {
	a := 2
	b := 4
	defer Add(a, b)
	a = 3
	defer Add(a, b)
	b = 5
}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func TestAA(t *testing.T) {
	s := NewSlice()
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%#v", o.Quantity)
}

func TestAA1(t *testing.T) {
	//var orange Orange
	//orange.Increase(10)
	//orange.Decrease(5)
	//fmt.Println(orange)

	//slice := []int{0, 1, 2, 3}
	//m := make(map[int]*int)
	//
	//for key, val := range slice {
	//	m[key] = &val
	//	fmt.Println(key, "--->", val)
	//}
	//
	//for k, v := range m {
	//	fmt.Println("v:", v)
	//	fmt.Println(k, "->", *v)
	//}

	//s := make([]int, 0)
	//	//s = append(s, 1, 2, 3)
	//	//fmt.Println(s)

	//list := make([]int, 0)
	//list = append(list, 1)
	//fmt.Println(list)

	//s1 := []int{1, 2, 3}
	//s2 := []int{4, 5}
	//s1 = append(s1, s2)
	//fmt.Println(s1)

	var (
		size     = 1024
		max_size = size * 2
	)
	fmt.Println(size, max_size)

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}

	//if sm1.m == sm2.m {
	//	fmt.Println("sm1 == sm2")
	//}
}

func TestPrintLog(t *testing.T) {
	var ch = make(chan string, 100)
	var wg sync.WaitGroup
	wg.Add(2)
	go send(ch, &wg)
	go receive(ch, &wg)

	wg.Wait()
	fmt.Println("over")

	defer close(ch)
}

func send(ch chan string, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		ch <- "消息：" + strconv.Itoa(i)
	}
	wg.Done()
}

func receive(ch chan string, wg *sync.WaitGroup) {
	var loglist [100]string
	y := 0
	for {
		if y >= 10000 {
			fmt.Println("y >= 10000", y)
			wg.Done()
			return
		}
		for i := 0; i < 100; i++ {
			select {
			case loglist[i] = <-ch:
				y++
			case <-time.After(time.Second * 1):
				return
			}
		}
		fmt.Println(loglist)
	}
}
