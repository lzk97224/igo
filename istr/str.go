package istr

import "unsafe"

// ToBytes 字符串转字节切片
func ToBytes(src string) []byte {
	pointer := (*[2]uintptr)(unsafe.Pointer(&src))
	tmp := [3]uintptr{pointer[0], pointer[1], pointer[1]}
	return *(*[]byte)(unsafe.Pointer(&tmp))
}
