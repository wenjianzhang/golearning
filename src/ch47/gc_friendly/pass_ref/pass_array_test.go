package pass_ref

import "testing"

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumOfElems]Content) int {
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content

	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWithValue(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}
