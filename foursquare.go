package foursquarego

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	baseURL = "https://api.foursquare.com/v2/"
	version = "20150813"
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
