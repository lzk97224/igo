package imix

import (
	"fmt"
	"log"
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
func SimpleRetry[T any](do func() (data T, err error), maxTimes int) (result T, err error) {
	for i := 0; i < maxTimes; i++ {
		log.Printf("do,times :%v\n", i+1)
		result, err = do()
		if err == nil {
			return result, nil
		}
	}
	var t T
	return t, fmt.Errorf("SimpleRetry all fail:%w", err)
}

// If 三目运算
func If[T any](con bool, one, another T) T {
	if con {
		return one
	} else {
		return another
	}
}
