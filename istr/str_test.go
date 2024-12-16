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

func TestIsChinese(t *testing.T) {
	fmt.Println(IsContainsChinese("jj12"))
	fmt.Println(IsContainsChinese("j我j12"))
	fmt.Println(IsContainsChinese("j孓j12"))
	fmt.Println(IsContainsChinese("j亻j12"))
}

func TestContainsAny(t *testing.T) {
	fmt.Println(ContainsAny("abc", "a"))
	fmt.Println(ContainsAny("abc", "d", "e"))
	fmt.Println(ContainsAny("abc", ""))
}

func Test_trunc(t *testing.T) {
	fmt.Println(Trunc("12345", 0, 3))
	fmt.Println(Trunc("12345", 1, 3))
	fmt.Println(Trunc("12345", 2, 3))
	fmt.Println(Trunc("12345", 3, 3))
	fmt.Println(Trunc("12345", 4, 3))
	fmt.Println(Trunc("12345", 5, 3))

	fmt.Println(TruncLeft("12345", 3))
	fmt.Println(TruncLeft("12345", 5))
	fmt.Println(TruncLeft("12345", 6))

	fmt.Println(TruncRight("12345", 3))
	fmt.Println(TruncRight("12345", 5))
	fmt.Println(TruncRight("12345", 6))
}
