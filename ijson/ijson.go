package ijson

import (
	"bytes"
	"encoding/json"
	"github.com/lzk97224/igo/ibytes"
	"github.com/lzk97224/igo/istr"
)

// StringToObj json字符串转对象
func StringToObj[T any](jsonStr string) (T, error) {
	return BytesToObj[T](istr.ToBytes(jsonStr))
}

// BytesToObj json字节切片转对象
func BytesToObj[T any](jsonByte []byte) (T, error) {
	decoder := json.NewDecoder(bytes.NewBuffer(jsonByte))
	decoder.UseNumber()
	var t T
	err := decoder.Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}

// ObjToString 对象转成json字符串
func ObjToString(v interface{}) string {
	return ibytes.ToString(ObjToBytes(v))
}

// ObjToBytes 对象转json字节切片
func ObjToBytes(v interface{}) []byte {
	jsonStu, err := json.Marshal(v)
	if err != nil {
		return []byte{'n', 'u', 'l', 'l'}
	}
	return jsonStu
}

// Convert  对象转换
func Convert[T any](src interface{}) (T, error) {
	return BytesToObj[T](ObjToBytes(src))
}
