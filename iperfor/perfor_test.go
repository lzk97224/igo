package iperfor

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	DotBegin("aaa")
	time.Sleep(time.Second * 2)
	DotEnd("aaa")

	DotBegin("bbb")
	time.Sleep(time.Second * 1)
	DotEnd("bbb")
	fmt.Println(Print())

	fmt.Println(GetDot("bbb"))

	for i := 0; i < 100; i++ {
		a := i
		go func() {
			DotBegin(fmt.Sprintf("%d", a))
		}()
	}
	time.Sleep(time.Second)

}
