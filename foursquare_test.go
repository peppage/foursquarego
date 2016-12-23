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

const clientSecret = "cs"
const clientID = "ci"

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

func assertQueryNoUser(t *testing.T, expected map[string]string, req *http.Request) {
	expected["v"] = version
	expected["m"] = "foursquare"
	expected["client_id"] = clientID
	expected["client_secret"] = clientSecret

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

func TestRateLimit(t *testing.T) {
	resp := http.Response{
		Header: make(http.Header),
	}

	resp.Header.Add(headerRateLimit, "5000")
	resp.Header.Add(headerRatePath, "/v2/venues/X")
	resp.Header.Add(headerRateRemaining, "4999")

	rl := RateLimitData(&resp)

	assert.Equal(t, 5000, rl.Limit)
	assert.Equal(t, "/v2/venues/X", rl.Path)
	assert.Equal(t, 4999, rl.Remaining)
}
