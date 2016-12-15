package foursquarego

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

type RewriteTransport struct {
	Transport http.RoundTripper
}

func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	assert.Equal(t, expectedMethod, req.Method)
}

func assertQuery(t *testing.T, expected map[string]string, req *http.Request) {
	expected["v"] = version
	queryValues := req.URL.Query()
	expectedValues := url.Values{}

	for key, value := range expected {
		expectedValues.Add(key, value)
	}
	assert.Equal(t, expectedValues, queryValues)
}

func getTestFile(path string) ([]byte, error) {
	// Open file with sample json
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}
