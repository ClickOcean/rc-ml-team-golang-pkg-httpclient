package httpclient

import (
	"context"
	"net/http"
)

type Client interface {
	GET(
		ctx context.Context,
		param RequestParams,
	) (*http.Response, error)
	PUT(
		ctx context.Context,
		param RequestParams,
	) (*http.Response, error)
	POST(
		ctx context.Context,
		param RequestParams,
	) (*http.Response, error)
}
