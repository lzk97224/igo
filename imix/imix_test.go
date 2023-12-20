package imix

import (
	"net/http"
	"testing"
)

func TestSimpleRetry(t *testing.T) {
	result, err := SimpleRetry(func() (data *http.Response, ok bool) {
		resp, err := http.Get("http://google.com")
		return resp, err == nil
	}, 3)
	if err != nil {
		t.Log(result, err)
	}

	result, err = SimpleRetry(func() (data *http.Response, ok bool) {
		resp, err := http.Get("http://baidu.com")
		return resp, err == nil
	}, 3)
	if err != nil {
		t.Log(result, err)
	}
}
