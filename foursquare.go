package foursquarego

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dghubble/sling"
)

const (
	baseURL             = "https://api.foursquare.com/v2/"
	version             = "20161213"
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "x-RateLimit-Remaining"
	headerRatePath      = "X-RateLimit-Path"
)

// Client is a Foursquare client for making Foursquare API requests.
type Client struct {
	sling *sling.Sling

	// Services used for talking to different parts of the API
	Venues *VenueService
}

// NewClient returns a new Client.
func NewClient(httpClient *http.Client, mode, clientID, clientSecret, accessToken string) *Client {
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
// https://developer.foursquare.com/docs/api/getting-started#6-make-your-first-api-call
type Response struct {
	Meta          Meta            `json:"meta"`
	Notifications []Notification  `json:"notifications"`
	Response      json.RawMessage `json:"response"`
}

// Meta contains request information and error details
// https://developer.foursquare.com/docs/api/troubleshooting/errors
type Meta struct {
	Code        int    `json:"code"`
	ErrorType   string `json:"errorType"`
	ErrorDetail string `json:"errorDetail"`
	RequestID   string `json:"requestId"`
}

// Notification comes with all responses.
// https://developer.foursquare.com/docs/responses/notifications
type Notification struct {
	Type string  `json:"type"`
	Item Omitted `json:"item"`
}

// Group contains the default fields in a group. A lot of responses
// share these fields.
type Group struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// ID is a simple struct with just an id. VenuePage is an example.
type ID struct {
	ID string `json:"id"`
}

// Count is a simple struct with just a count. Followers and todo are examples
// that only have a count.
type Count struct {
	Count int `json:"count"`
}

// Omitted is for fields that do not have a known datastructure. If you find
// an example where this field is used please let me know. You will need to handle
// this in your application until that time.
type Omitted interface{}

// BoolAsAnInt is a bool that needs to be an int when transferred to an endpoint
type BoolAsAnInt int

// Option available for BoolAsAnInt
const (
	True = BoolAsAnInt(1)
)

// RateLimit is a struct of foursquare ratelimit data
type RateLimit struct {
	Limit     int
	Path      string
	Remaining int
}

// ParseRate is a helper function to get all the Rate info
func ParseRate(resp *http.Response) *RateLimit {
	limit := resp.Header.Get(headerRateLimit)
	path := resp.Header.Get(headerRatePath)
	remain := resp.Header.Get(headerRateRemaining)

	l, _ := strconv.Atoi(limit)
	r, _ := strconv.Atoi(remain)

	return &RateLimit{
		Limit:     l,
		Path:      path,
		Remaining: r,
	}
}
