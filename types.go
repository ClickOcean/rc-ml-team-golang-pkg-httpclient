package httpclient

type RetryConfig struct {
	RetryCount    int
	MaxBackoffSec int
	AttemptCodes  []int
}

type RequestParams struct {
	URL     string
	Headers map[string]string
	//Body set the request Body, accepts string, []byte, io.Reader, map and struct.
	Body any
	//SetErrorResult set the result that response body will be
	//unmarshalled to if no error occurs and Response.ResultState() returns ErrorState,
	//by default it requires HTTP status `code >= 400`
	ErrorResult any
	//SetSuccessResult set the result that response Body will be
	//unmarshalled to if no error occurs and Response.ResultState() returns SuccessState,
	//by default it requires HTTP status `code >= 200 && code <= 299`
	SuccessResult any
}
