package iperfor

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	DotBegin("aaa")
	time.Sleep(time.Second * 2)
	DotEnd("aaa")
	Print()
}
