package imix

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// SimpleRetry 一个简单的重试工具
//
// 参数：
//
//	do:处理函数， 返回 data是函数返回数据，也是SimpleRetry最终返回的数据；ok 表示是否处理成功，成功则结束重试
//	maxTimes: 最大重试次数
//
// 返回值：
//
//	result: 对应do函数的返回值
//	err: 代表最终结果是否成功
func SimpleRetryWithDelay[T any](do func() (data T, err error), maxTimes int, delay time.Duration) (result T, err error) {
	for i := 0; i < maxTimes; i++ {
		log.Printf("do,times :%v\n", i+1)
		result, err = do()
		if err == nil {
			return result, nil
		}
		time.Sleep(delay)
	}
	var t T
	return t, fmt.Errorf("SimpleRetry all fail:%w", err)
}

func SimpleRetry[T any](do func() (data T, err error), maxTimes int) (result T, err error) {
	return SimpleRetryWithDelay(do, maxTimes, time.Millisecond*100)
}

// If 三目运算
func If[T any](con bool, one, another T) T {
	if con {
		return one
	} else {
		return another
	}
}

func IfDelay[T any](con bool, one, another func() T) T {
	if con {
		return one()
	} else {
		return another()
	}
}

func DeepCopy[T any](src T) (T, error) {
	var buf bytes.Buffer
	var dst T
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return dst, err
	}

	err := gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(&dst)

	return dst, err
}
