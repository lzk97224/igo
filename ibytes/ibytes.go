package ibytes

import "unsafe"

func ToString(src []byte) string {
	return *(*string)(unsafe.Pointer(&src))
}
