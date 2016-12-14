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
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	Contact          Contact    `json:"contact"`
	Location         Location   `json:"location"`
	CanonicalURL     string     `json:"canonicalUrl"`
	Categories       []Category `json:"categories"`
	Verified         bool       `json:"verified"`
	Stats            Stats      `json:"stats"`
	URL              string     `json:"url"`
	Price            Price      `json:"price"`
	HasMenu          bool       `json:"hasMenu"`
	Likes            Likes      `json:"likes"`
	Like             bool       `json:"like"`
	Dislike          bool       `json:"dislike"`
	Ok               bool       `json:"ok"`
	Rating           float64    `json:"rating"`
	RatingColor      string     `json:"ratingColor"`
	RatingSignals    int        `json:"ratingSignals"`
	Menu             Menu       `json:"menu"`
	AllowMenuURLEdit bool       `json:"allowMenuUrlEdit"`
	/*
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

type Contact struct {
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
}

type Location struct {
	Address          string           `json:"address"`
	CrossStreet      string           `json:"crossStreet"`
	Lat              float64          `json:"lat"`
	Lng              float64          `json:"lng"`
	LabeledLatLngs   []LabeledLatLngs `json:"labeledLatLngs"`
	PostalCode       string           `json:"postalCode"`
	Cc               string           `json:"cc"`
	City             string           `json:"city"`
	State            string           `json:"state"`
	Country          string           `json:"country"`
	FormattedAddress []string         `json:"formattedAddress"`
	IsFuzzed         bool             `json:"isFuzzed"`
	Distance         int              `json:"distance"`
}

type LabeledLatLngs struct {
	Label string  `json"label"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

type Category struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PluralName string `json:"pluralName"`
	ShortName  string `json:"shortName"`
	Icon       Icon   `json:"icon"`
	Primary    bool   `json:"primary"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Stats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
	VisitsCount   int `json:"visitsCount"`
}

type Price struct {
	Tier     int    `json:"tier"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
}

type Likes struct {
	Count int `json:"count"`
	//Groups Omit `json:"-"` //TODO: take care fo this later
	//Items   []User `json:"items"` // I think this is under groups?
	Summary string `json:"summary"`
}

type Menu struct {
	Type      string `json:"type"`
	Label     string `json:"label"`
	Anchor    string `json:"anchor"`
	URL       string `json:"url"`
	MobileURL string `json:"mobileUrl"`
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
