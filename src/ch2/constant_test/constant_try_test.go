package constart_test

import (
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable  = 1 << iota //0001
	Writable              //0010
	Exectable             //0011
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

}

func TestConstantTry1(t *testing.T) {
	a := 6
	t.Log(uint8(a), uint(Readable), uint(Writable), uint(Exectable))
	c := a & Readable
	t.Logf("第一行 - c 的值为 %d\n", c)
	c = a & Writable
	t.Logf("第一行 - c 的值为 %d\n", c)
	c = a & Exectable
	t.Logf("第一行 - c 的值为 %d\n", c)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Exectable == Exectable)
}
