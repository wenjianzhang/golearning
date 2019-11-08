package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c) // c长度不一样，这里直接就会报编译错误
	t.Log(a == d)
}

const (
	Readable  = 1 << iota //0001
	Writable              //0010
	Exectable             //0011
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	a = a &^ Exectable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Exectable == Exectable)
}
