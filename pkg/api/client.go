package api

import (
	"github.com/omniful/go_commons/httpclient"
	"time"
)

func NewClient(host string, opts ...httpclient.Option) httpclient.Client {
	os := httpclient.Options{
		httpclient.WithLogConfig(httpclient.LogConfig{
			LogRequest:  true,
			LogResponse: true,
		}),
		httpclient.WithDeadline(time.Minute),
	}
	os = append(os, opts...)
	return httpclient.New(host, os...)
}
