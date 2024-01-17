package imix

import (
	"fmt"
	"github.com/lzk97224/igo/istr"
	"net/http"
	"testing"
	"time"
)

func TestSimpleRetry(t *testing.T) {
	http.DefaultClient.Timeout = time.Second * 2
	result, err := SimpleRetry(func() (data *http.Response, err error) {
		resp, err := http.Get("http://google.com")
		return resp, err
	}, 3)
	if err != nil {
		t.Log(result, err)
	}

	result, err = SimpleRetry(func() (data *http.Response, err error) {
		resp, err := http.Get("http://baidu.com")
		return resp, err
	}, 3)
	if err != nil {
		t.Log(result, err)
	}

	_, err = SimpleRetry(
		func() (data string, err error) {
			resp, err := "", fmt.Errorf("sdf")
			if err != nil {
				return resp, err
			}
			if !istr.IsContainsChinese(resp) {
				return resp, fmt.Errorf("翻译结果包含汉字:%v", resp)
			}
			return resp, nil
		},
		3)
	if err != nil {
		fmt.Println(fmt.Errorf("翻译提示词失败,%w", err))
	}

	_, err = SimpleRetry(
		func() (data string, err error) {
			resp, err := "我", nil
			if err != nil {
				return resp, err
			}
			if istr.IsContainsChinese(resp) {
				return resp, fmt.Errorf("翻译结果包含汉字:%v", resp)
			}
			return resp, nil
		},
		3)
	if err != nil {
		fmt.Println(fmt.Errorf("翻译提示词失败,%w", err))
	}
}
