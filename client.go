package httpclient

import (
	"context"
	"net/http"
	"time"

	"github.com/imroc/req/v3"
	"golang.org/x/exp/slices"
)

type client struct {
	client *req.Client
}

func New() *client {
	return &client{
		req.C(),
	}
}

func (c *client) WithRetry(cfg RetryConfig) *client {
	c.client.
		SetCommonRetryCount(cfg.RetryCount).
		SetCommonRetryBackoffInterval(100*time.Millisecond, time.Duration(cfg.MaxBackoffSec)*time.Second).
		AddCommonRetryCondition(func(resp *req.Response, err error) bool {
			return err != nil || slices.Contains(cfg.AttemptCodes, resp.StatusCode)
		})

	return c
}

func (c *client) WithTimeout(sec int) *client {
	c.client.SetTimeout(time.Duration(sec) * time.Second)
	return c
}

func (c client) GET(
	ctx context.Context,
	param RequestParams,
) (*http.Response, error) {
	r := c.client.Get(param.URL).
		SetSuccessResult(param.SuccessResult).
		SetErrorResult(param.ErrorResult).
		SetHeaders(param.Headers).
		Do(ctx)

	return r.Response, r.Err
}

func (c client) PUT(
	ctx context.Context,
	param RequestParams,
) (*http.Response, error) {
	r := c.client.Put(param.URL).
		SetSuccessResult(param.SuccessResult).
		SetErrorResult(param.ErrorResult).
		SetHeaders(param.Headers).
		SetBody(param.Body).
		Do(ctx)

	return r.Response, r.Err
}

func (c client) POST(
	ctx context.Context,
	param RequestParams,
) (*http.Response, error) {
	r := c.client.Post(param.URL).
		SetSuccessResult(param.SuccessResult).
		SetErrorResult(param.ErrorResult).
		SetHeaders(param.Headers).
		SetBody(param.Body).
		Do(ctx)

	return r.Response, r.Err
}
