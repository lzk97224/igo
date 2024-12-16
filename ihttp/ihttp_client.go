package ihttp

import (
	"context"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"time"
)

// NewHttpClient 创建新的http客户端,并且指定写超时和读超时
func NewHttpClient(writeTimeout time.Duration, readTimeout time.Duration) *resty.Client {
	c := resty.New()
	transport := http.Transport{
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			c, err := net.DialTimeout(network, address, time.Second*5)
			if err != nil {
				return nil, err
			}
			err = c.SetWriteDeadline(time.Now().Add(time.Second * writeTimeout))
			if err != nil {
				return nil, err
			}
			err = c.SetReadDeadline(time.Now().Add(time.Second * (writeTimeout + readTimeout)))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	c.SetTransport(&transport)
	return c
}
