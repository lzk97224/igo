package ihttp

import (
	"fmt"
	"testing"
)

func TestGetWithTimeout(t *testing.T) {
	result := map[string]any{}
	timeout, m, err := Get("https://www.baidu.com", nil, nil, &result)
	fmt.Println(timeout, m, err)
}
