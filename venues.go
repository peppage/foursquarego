package foursquarego

import (
	"encoding/json"
	"net/http"

	"github.com/dghubble/sling"
)

// Need this since the responses all have subitems.
type venueResp struct {
	Venue Venue `json:"venue"`
}

// Venue represents a foursquare Venue.
// https://developer.foursquare.com/docs/responses/venue
type Venue struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	/*Contact       Contact    `json:"contact"`
	Location      Location   `json:"location"`
	CanonicalUrl  string     `json:"canonicalUrl"`
	Categories    []Category `json:"categories"`
	Verified      bool       `json:"verified"`
	Stats         Stats      `json:"stats"`
	Url           string     `json:"url"`
	Price         Price      `json:"price"`
	HasMenu       bool       `json:"hasMenu"`
	Likes         Likes      `json:"likes"`
	Like          bool       `json:"like"`
	Dislike       bool       `json:"dislike"`
	Ok            bool       `json:"ok"`
	Rating        float32    `json:"rating"`
	RatingColor   string     `json:"ratingColor"`
	RatingSignals int        `json:"ratingSignals"`
	Menu          ShortMenu  `json:"menu"`
	Specials      Specials   `json:"specials"`
	Photos        Photos     `json:"photos"`
	Reasons       Reasons    `json:"reasons"`
	CreatedAt     int        `json:"createdAt"`
	Tips          Tips       `json:"tips"`
	Tags          []string   `json:"tags"`
	ShortUrl      string     `json:"shortUrl"`
	TimeZone      string     `json:"timeZone"`
	Listed        Listed     `json:"listed"`
	Phrases       []Phrase   `json:"phrases"`
	Hours         HoursVenue `json:"hours"`
	Popular       Popular    `json:"popular"`
	PageUpdates   Omit       `json:"-"`
	Inbox         Omit       `json:"-"`
	VenueChains   Omit       `json:"-"`
	Attributes    Attributes `json:"attributes"`
	BestPhoto     Photo      `json:"bestPhoto"`*/
}

// VenueService provies a method for accessing Foursquare venue endpoints
type VenueService struct {
	sling *sling.Sling
}

func newVenueService(sling *sling.Sling) *VenueService {
	return &VenueService{
		sling: sling.Path("venues/"),
	}
}

func (s *VenueService) Details(id string) (*Venue, *http.Response, error) {
	response := new(Response)
	venue := new(venueResp)

	resp, err := s.sling.New().Get(id).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venue)
	}
	return &venue.Venue, resp, relevantError(err, *response)
}
