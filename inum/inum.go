package inum

import "github.com/lzk97224/igo/imix"

type Number interface {
	int | int64 | int32 | int16 | int8 | uint | uint64 | uint32 | uint16 | uint8 | float32 | float64
}

// Min 求最大值
func Min[T Number](one, anoter T) T {
	return imix.If(one < anoter, one, anoter)
}

// Max 求最大值
func Max[T Number](one, anoter T) T {
	return imix.If(one > anoter, one, anoter)
}
