package pipe_filter

import (
	"fmt"
	"testing"
)

func TestStraightPipeline(t *testing.T) {
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeline("p1", spliter, converter, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}

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

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1.m == sm2.m {
		fmt.Println("sm1 == sm2")
	}
}
