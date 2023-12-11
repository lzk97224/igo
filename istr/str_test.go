package istr

import (
	"fmt"
	"testing"
)

func TestStringToByteSlice(t *testing.T) {
	fmt.Printf("%s\n", ToBytes("水电费健康"))
	fmt.Printf("%s\n", ToBytes("123"))
	fmt.Printf("%s\n", ToBytes(""))

	s := append(ToBytes("123"), 's')
	fmt.Printf("%s\n", s)
	ss := ToBytes("123")
	ss[1] = 's'
	fmt.Printf("%s\n", ss)
}
