package foursquarego

import (
	"encoding/json"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	baseURL = "https://api.foursquare.com/v2/"
	version = "20161213"
	//MODE    = "foursquare"
)

// Client is a Foursquare client for making Foursquare API requests.
type Client struct {
	sling  *sling.Sling
	Venues *VenueService
}

// NewClient returns a new Client.
func NewClient(mode, clientID, clientSecret, accessToken string) *Client {
	httpClient := http.DefaultClient
	return newClient(httpClient, mode, clientID, clientSecret, accessToken)
}

func newClient(httpClient *http.Client, mode, clientID, clientSecret, accessToken string) *Client {
	b := sling.New().Client(httpClient).Base(baseURL)
	b.QueryStruct(struct {
		V            string `url:"v"`
		M            string `url:"m"`
		ClientID     string `url:"client_id"`
		ClientSecret string `url:"client_secret,omitempty"`
		AccessToken  string `url:"access_token,omitempty"`
	}{
		V:            version,
		M:            mode,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AccessToken:  accessToken,
	})

	return &Client{
		sling:  b,
		Venues: newVenueService(b.New()),
	}
}

// RawRequest allows you to make any request you want. This will automatically add
// the client/user tokens. Gives back exactly the response from foursquare.
func (c *Client) RawRequest(url string) (*Response, *http.Response, error) {
	response := new(Response)
	resp, err := c.sling.New().Get(url).Receive(response, response)
	return response, resp, relevantError(err, *response)
}

// Response is a typical foursquare response
// https://developer.foursquare.com/overview/responses
type Response struct {
	Meta          Meta            `json:"meta"`
	Notifications []Notification  `json:"notifications"`
	Response      json.RawMessage `json:"response"`
}

type Meta struct {
	Code        int    `json:"code"`
	ErrorType   string `json:"errorType"`
	ErrorDetail string `json:"errorDetail"`
	RequestID   string `json:"requestId"`
}

type Notification struct {
	Type string          `json:"type"`
	Item json.RawMessage `json:"item"`
}

type Group struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Count struct {
	Count int `json:"count"`
}

// Omitted is for fields that do not have a known datastructure. If you find
// an example where this field is used please let me know. You will need to handle
// this in your application until that time.
type Omitted interface{}
