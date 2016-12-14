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
func NewClient(httpClient *http.Client) *Client {
	b := sling.New().Client(httpClient).Base(baseURL)
	return &Client{
		sling:  b,
		Venues: newVenueService(b.New()),
	}
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
