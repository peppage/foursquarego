package foursquarego

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FoursquareApi struct {
	clientID     string
	clientSecret string
	queryQueue   chan query
	HttpClient   *http.Client
}

type query struct {
	url         string
	form        url.Values
	data        *foursquareResponse
	method      int
	response_ch chan response
}

type response struct {
	data interface{}
	err  error
}

type Omit struct{}

const (
	API_URL = "https://api.foursquare.com/v2/"
	VERSION = "20150813"
	MODE    = "foursquare"
	_GET    = iota
	_POST   = iota
)

func NewFoursquareApi(clientID string, clientSecret string) *FoursquareApi {
	queue := make(chan query)
	a := &FoursquareApi{
		clientID:     clientID,
		clientSecret: clientSecret,
		queryQueue:   queue,
		HttpClient:   http.DefaultClient,
	}
	go a.throttledQuery()
	return a
}

func (a *FoursquareApi) apiGet(urlStr string, form url.Values, data *foursquareResponse) error {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}

	req.URL.RawQuery = form.Encode()
	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResp apiResponse
	json.Unmarshal(contents, &apiResp)
	if resp.StatusCode != 200 || apiResp.Meta.Code != 200 {
		return newApiError(resp, apiResp.Meta)
	}

	*data = apiResp.Response
	return nil
}

func cleanValues(v url.Values) url.Values {
	if v == nil {
		return url.Values{}
	}
	return v
}

func (a *FoursquareApi) execQuery(urlStr string, form url.Values, data *foursquareResponse, method int) error {
	form.Set("v", VERSION)
	form.Set("m", MODE)
	if form.Get("oauth_token") == "" {
		form.Set("client_id", a.clientID)
		form.Set("client_secret", a.clientSecret)
	}
	switch method {
	case _GET:
		return a.apiGet(urlStr, form, data)
	default:
		return fmt.Errorf("HTTP method not supported")
	}
	return errors.New("ack")
}

func (a *FoursquareApi) throttledQuery() {
	for q := range a.queryQueue {
		url := q.url
		form := q.form
		data := q.data
		method := q.method

		response_ch := q.response_ch

		err := a.execQuery(url, form, data, method)
		response_ch <- response{data, err}
	}
}

func (a *FoursquareApi) Close() {
	close(a.queryQueue)
}
