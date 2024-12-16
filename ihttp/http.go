package ihttp

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

var defaultClient = NewHttpClient(5, 5)

func checkResult[T any](res *resty.Response, err error, result T) (*resty.Response, T, error) {
	if err != nil {
		return res, result, err
	}
	if !res.IsSuccess() {
		return res, result, fmt.Errorf("request fail code:%v,msg:%v", res.StatusCode(), res.String())
	}
	if res.RawResponse.Header.Get("Content-Type") == "text/plain" {
		err = json.Unmarshal(res.Body(), result)
	}
	return res, result, nil
}

func createReqWithoutBody[T any](c *resty.Client, headers map[string]string, queryParams map[string]string, result T) *resty.Request {
	r := c.R()
	r.SetHeaders(headers)
	r.SetQueryParams(queryParams)
	r.SetResult(result)
	return r
}

func createReqWithBody[T any](c *resty.Client, headers map[string]string, body any, result T) *resty.Request {
	r := c.R()
	r.SetHeaders(headers)
	r.SetBody(body)
	r.SetResult(result)
	return r
}

func GetWithClient[T any](c *resty.Client, url string, headers map[string]string, queryParams map[string]string, result T) (*resty.Response, T, error) {
	r := createReqWithoutBody(c, headers, queryParams, result)
	res, err := r.Get(url)
	return checkResult(res, err, result)
}

func PostWithClient[T any](c *resty.Client, url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	r := createReqWithBody(c, headers, body, result)
	res, err := r.Post(url)
	return checkResult(res, err, result)
}
func PutWithClient[T any](c *resty.Client, url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	r := createReqWithBody(c, headers, body, result)
	res, err := r.Put(url)
	return checkResult(res, err, result)
}

func DelWithClient[T any](c *resty.Client, url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	r := createReqWithBody(c, headers, body, result)
	res, err := r.Delete(url)
	return checkResult(res, err, result)
}

func Get[T any](url string, headers map[string]string, queryParams map[string]string, result T) (*resty.Response, T, error) {
	return GetWithClient(defaultClient, url, headers, queryParams, result)
}

func Post[T any](url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	return PostWithClient(defaultClient, url, headers, body, result)
}
func Put[T any](url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	return PutWithClient(defaultClient, url, headers, body, result)
}

func Del[T any](url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	return DelWithClient(defaultClient, url, headers, body, result)
}
