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
	Open    = 1 << iota //0001
	Close               //0010
	Pending             //0011
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

}

func TestConstantTry1(t *testing.T) {
	a := 6
	t.Log(uint8(a), uint(Open), uint(Close), uint(Pending))
	c := a & Open
	t.Logf("第一行 - c 的值为 %d\n", c)
	c = a & Close
	t.Logf("第一行 - c 的值为 %d\n", c)
	c = a & Pending
	t.Logf("第一行 - c 的值为 %d\n", c)
	t.Log(a&Open == Open, a&Close == Close, a&Pending == Pending)
}
