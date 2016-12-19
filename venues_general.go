package foursquarego

import (
	"encoding/json"
	"errors"
	"net/http"
)

type categoriesResp struct {
	Categories []Category `json:"categories"`
}

// Categories returns a hierarchical list of categories applied to venues.
// https://developer.foursquare.com/docs/venues/categories
func (s *VenueService) Categories() ([]Category, *http.Response, error) {
	cats := new(categoriesResp)
	response := new(Response)
	resp, err := s.sling.New().Get("categories").Receive(response, response)

	if err == nil {
		json.Unmarshal(response.Response, cats)
	}

	return cats.Categories, resp, relevantError(err, *response)
}

// SearchIntentParam represents the options available for Intent on the search endpoint
type SearchIntentParam string

// Options that are valid for the Intent on the search endpoint
const (
	CheckinIntent = SearchIntentParam("checkin")
	BrowseIntent  = SearchIntentParam("browse")
	GlobalIntent  = SearchIntentParam("global")
	MatchIntent   = SearchIntentParam("match")
)

// VenueSearchParams are the parameters for the VenueService.Search
type VenueSearchParams struct {
	LatLong          string            `url:"ll,omitempty"`
	Near             string            `url:"near,omitempty"`
	LatLongAccuracy  int               `url:"llAcc,omitempty"`
	Altitude         int               `url:"alt,omitempty"`
	AltitudeAccuracy int               `url:"altAcc,omitempty"`
	Query            string            `url:"query,omitempty"`
	Limit            int               `url:"limit,omitempty"`
	Intent           SearchIntentParam `url:"intent,omitempty"`
	Radius           int               `url:"raidus,omitempty"`
	Sw               string            `url:"sw,omitempty"`
	Ne               string            `url:"ne,omitempty"`
	CategoryID       []string          `url:"categoryId,omitempty"`
	URL              string            `url:"url,omitempty"`
	ProviderID       string            `url:"providerId,omitempty"`
	LinkedID         int               `url:"linkedId,omitempty"`
}

type venueSearchResp struct {
	Venues []Venue `json:"venues"`
}

// Search returns a list of venues near the current location, optionally matching a search term.
// https://developer.foursquare.com/docs/venues/search
func (s *VenueService) Search(params *VenueSearchParams) ([]Venue, *http.Response, error) {
	if params.LatLong == "" && params.Near == "" {
		return nil, nil, errors.New("LatLong or Near are required")
	}

	venues := new(venueSearchResp)
	response := new(Response)

	resp, err := s.sling.New().Get("search").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venues)
	}

	return venues.Venues, resp, relevantError(err, *response)
}

// VenueSuggestParams are the parementers for the VenueService.SuggestCompletion
type VenueSuggestParams struct {
	LatLong          string `url:"ll,omitempty"`
	Near             string `url:"near,omitempty"`
	LatLongAccuracy  int    `url:"llAcc,omitempty"`
	Altitude         int    `url:"alt,omitempty"`
	AltitudeAccuracy int    `url:"altAcc,omitempty"`
	Query            string `url:"query,omitempty"`
	Limit            int    `url:"limit,omitempty"`
	Radius           int    `url:"raidus,omitempty"`
	Sw               string `url:"sw,omitempty"`
	Ne               string `url:"ne,omitempty"`
}

type MiniVenue struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Location Location   `json:"location"`
	Category []Category `json:"categories"`
	HasPerk  bool       `json:"hasPerk"`
}

type venueSuggestResp struct {
	MiniVenues []MiniVenue `json:"minivenues"`
}

// SuggestCompletion returns a list of mini-venues partially matching the search term, near the location.
// https://developer.foursquare.com/docs/venues/suggestcompletion
func (s *VenueService) SuggestCompletion(params *VenueSuggestParams) ([]MiniVenue, *http.Response, error) {
	if params.LatLong == "" && params.Near == "" {
		return nil, nil, errors.New("LatLong or Near are required")
	}

	if params.Query == "" {
		return nil, nil, errors.New("Query is required")
	}

	venues := new(venueSuggestResp)
	response := new(Response)

	resp, err := s.sling.New().Get("suggestCompletion").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venues)
	}

	return venues.MiniVenues, resp, relevantError(err, *response)
}

// VenueTrendingParams are the parameters for the VenueService.Trending endpoint
type VenueTrendingParams struct {
	LatLong string `url:"ll,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	Radius  int    `url:"raidus,omitempty"`
}

type venueTrendingResp struct {
	Venues []Venue `json:"venues"`
}

// Trending returns a list of venues near the current location with the most people currently checked in.
// https://developer.foursquare.com/docs/venues/trending
func (s *VenueService) Trending(params *VenueTrendingParams) ([]Venue, *http.Response, error) {
	if params.LatLong == "" {
		return nil, nil, errors.New("LatLong is required")
	}

	venues := new(venueTrendingResp)
	response := new(Response)

	resp, err := s.sling.New().Get("trending").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venues)
	}

	return venues.Venues, resp, relevantError(err, *response)
}
