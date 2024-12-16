package ibean

import (
	"fmt"
	"reflect"
)

func CopyBean(src any, dest any) error {
	srcVal := reflect.ValueOf(src)
	destVal := reflect.ValueOf(dest)

	// 确保 src 和 dest 都是指针类型
	if srcVal.Kind() != reflect.Ptr || destVal.Kind() != reflect.Ptr {
		return fmt.Errorf("both src and dest must be pointers to structs")
	}

	// 获取指向的值
	srcVal = srcVal.Elem()
	destVal = destVal.Elem()

	// 确保源和目标都是结构体
	if srcVal.Kind() != reflect.Struct || destVal.Kind() != reflect.Struct {
		return fmt.Errorf("both src and dest must be structs")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		destField := destVal.Field(i)

		// 只拷贝可导出的字段
		if srcField.CanInterface() && destField.CanSet() {
			destField.Set(srcField)
		}
	}

	return nil
}
