package ihttp

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

var defaultClient = resty.New()

func init() {
	defaultClient.SetTimeout(time.Second * 10)
}

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

func createReqWithoutBody[T any](headers map[string]string, queryParams map[string]string, result T) *resty.Request {
	r := defaultClient.R()
	r.SetHeaders(headers)
	r.SetQueryParams(queryParams)
	r.SetResult(result)
	return r
}

func createReqWithBody[T any](headers map[string]string, body any, result T) *resty.Request {
	r := defaultClient.R()
	r.SetHeaders(headers)
	r.SetBody(body)
	r.SetResult(result)
	return r
}

func Get[T any](url string, headers map[string]string, queryParams map[string]string, result T) (*resty.Response, T, error) {
	r := createReqWithoutBody(headers, queryParams, result)
	res, err := r.Get(url)
	return checkResult(res, err, result)
}

func Post[T any](url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	r := createReqWithBody(headers, body, result)
	res, err := r.Post(url)
	return checkResult(res, err, result)
}
func Put[T any](url string, headers map[string]string, body any, result T) (*resty.Response, T, error) {
	r := createReqWithBody(headers, body, result)
	res, err := r.Put(url)
	return checkResult(res, err, result)
}
