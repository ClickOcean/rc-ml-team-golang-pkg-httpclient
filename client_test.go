package httpclient_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"gitlab.ml.rc.dating.com/rc-ml-team/golang-pkg/httpclient"
)

type Suite struct {
	suite.Suite
	c       httpclient.Client
	server  *httptest.Server
	baseUrl string
}

func TestClient(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	s.c = httpclient.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte(`{"test":"data"}`))
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad_req"}`))
	})

	s.server = httptest.NewServer(mux)
	s.baseUrl = s.server.URL
}

func (s *Suite) TearDownSuite() {
	s.server.Close()
}

func (s *Suite) TestGet() {

	successResp := &struct {
		Test string `json:"test"`
	}{}

	param := httpclient.RequestParams{
		URL:           fmt.Sprintf("%s/get", s.baseUrl),
		SuccessResult: successResp,
	}
	resp, err := s.c.Get(context.TODO(), param)
	if s.NoError(err) {
		s.Equal(http.StatusOK, resp.StatusCode)
		s.Equal("data", successResp.Test)
	}
}

func (s *Suite) TestPut() {

	errResp := &struct {
		Err string `json:"error"`
	}{}

	param := httpclient.RequestParams{
		URL:         fmt.Sprintf("%s/put", s.baseUrl),
		ErrorResult: errResp,
	}
	resp, err := s.c.Put(context.TODO(), param)
	if s.NoError(err) {
		s.Equal(http.StatusBadRequest, resp.StatusCode)
		s.Equal("bad_req", errResp.Err)
	}
}
func (s *Suite) TestPost() {
	param := httpclient.RequestParams{
		URL: fmt.Sprintf("%s/post", s.baseUrl),
	}
	resp, err := s.c.Post(context.TODO(), param)
	if s.NoError(err) {
		s.Equal(http.StatusOK, resp.StatusCode)
	}
}
