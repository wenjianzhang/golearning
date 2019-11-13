package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the exoected is %d, the actual %d", inputs[i], expected[i], ret)
		}
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("start")
	t.Fatal("Fail")
	fmt.Println("End")
}

func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, expected[i], ret)
		if ret != expected[i] {
			t.Errorf("input is %d, the exoected is %d, the actual %d", inputs[i], expected[i], ret)
		}
	}
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	//与性能测试无关代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 测试代码
	}
	b.StopTimer()
	//与性能测试无关代码
}
