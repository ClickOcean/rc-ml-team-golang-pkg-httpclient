package httpclient

import (
	"context"
	"net/http"
)

type Client interface {
	Get(
		ctx context.Context,
		param RequestParams,
	) (*http.Response, error)
	Put(
		ctx context.Context,
		param RequestParams,
	) (*http.Response, error)
	Post(
		ctx context.Context,
		param RequestParams,
	) (*http.Response, error)
}
