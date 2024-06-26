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

	result, err = SimpleRetryWithDelay(func() (data *http.Response, err error) {
		resp, err := http.Get("http://google.com")
		return resp, err
	}, 3, time.Second*2)
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

type People struct {
	Age    int
	Name   string
	Tag    []string
	Friend []People
	Enemy  []*People
}

func TestDeepCopy(t *testing.T) {
	i := 1
	pi := &i
	ic, err := DeepCopy(pi)
	fmt.Println(ic, err)

	ic2, err := DeepCopy(i)
	fmt.Println(ic2, err)

	p := &People{
		Age:    1,
		Name:   "lzk",
		Tag:    []string{"1", "2"},
		Friend: []People{{Age: 2, Name: "lzk1"}},
		Enemy:  []*People{{Age: 3, Name: "lzk2"}},
	}

	p1, err := DeepCopy(p)
	fmt.Println(p1, err)

	p2, err := DeepCopy(*p)
	fmt.Println(p2, err)

}
