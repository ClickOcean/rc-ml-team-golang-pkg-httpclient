package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	c       *client
	server  *httptest.Server
	baseUrl string
}

func TestClient(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	s.c = New()

	mux := http.NewServeMux()
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write([]byte(`{"test":"data"}`))
		w.WriteHeader(http.StatusOK)
	})

	s.server = httptest.NewServer(mux)
	s.baseUrl = s.server.URL
}

func (s *Suite) TearDownSuite() {
	s.server.Close()
}

func (s *Suite) TestDo() {

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/get", s.baseUrl), http.NoBody)
	resp, err := s.c.Do(req)
	if !s.NoError(err) {
		return
	}
	s.Equal(http.StatusOK, resp.StatusCode)

	var data map[string]any

	err = json.NewDecoder(resp.Body).Decode(&data)
	s.NoError(err)
	s.Equal("data", data["test"])

}
