package gorutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Millisecond * 50)
}
