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

// ExploreSectionParam represents the options that are available for the Section param on the Venue Explore endpoint
type ExploreSectionParam string

// Options that are valid for Section in the Venue Explore endpoint
const (
	SectionFood       = ExploreSectionParam("food")
	SectionDrink      = ExploreSectionParam("drink")
	SectionCoffee     = ExploreSectionParam("coffee")
	SectionShops      = ExploreSectionParam("shops")
	SectionArts       = ExploreSectionParam("arts")
	SectionOutdoors   = ExploreSectionParam("outdoors")
	SectionSights     = ExploreSectionParam("sights")
	SectionTrending   = ExploreSectionParam("trending")
	SectionSpecials   = ExploreSectionParam("specials")
	SectionNextVenues = ExploreSectionParam("nextVenues")
	SectionTopPicks   = ExploreSectionParam("topPicks")
)

// ExploreNoveltyParam represents the options that are available for the novelty param on the Venue Explore endpoint
type ExploreNoveltyParam string

// Options that are valid for Novelty in the Venue Explore endpoint
const (
	NoveltyNew = ExploreNoveltyParam("new")
	NoveltyOld = ExploreNoveltyParam("old")
)

// ExploreFriendVisitParam represents the options that are available for the FriendVisits on the Venue Explore endpoint
type ExploreFriendVisitParam string

// Options that are valid for FriendVisit in the Venue Explore endpoint
const (
	FriendVisited    = ExploreFriendVisitParam("visited")
	FriendNotVisited = ExploreFriendVisitParam("notvisited")
)

// ExploreTimeParam represents the options that are available for the Time & Dat params on the Venue Explore endpoint
type ExploreTimeParam string

// Option that are valid for Time & Day in the Venue Explore endpoint
const (
	TimeAny = ExploreTimeParam("any")
)

// BoolAsAnInt is a bool that needs to be an int when transferred to an endpoint
type BoolAsAnInt int

// Option available for BoolAsAnInt
const (
	True = BoolAsAnInt(1)
)

// VenueExploreParams are the parameters for the VenueService.Explore endpoint
type VenueExploreParams struct {
	LatLong          string                  `url:"ll,omitempty"`
	Near             string                  `url:"near,omitempty"`
	LatLongAccuracy  int                     `url:"llAcc,omitempty"`
	Altitude         int                     `url:"alt,omitempty"`
	AltitudeAccuracy int                     `url:"altAcc,omitempty"`
	Radius           int                     `url:"raidus,omitempty"`
	Section          ExploreSectionParam     `url:"section,omitempty"`
	Query            string                  `url:"query,omitempty"`
	Limit            int                     `url:"limit,omitempty"`
	Offset           int                     `url:"offset,omitempty"`
	Novelty          ExploreNoveltyParam     `url:"novelty,omitempty"`
	FriendVisits     ExploreFriendVisitParam `url:"friendVists,omitempty"`
	Time             ExploreTimeParam        `url:"time,omitempty"`
	Day              ExploreTimeParam        `url:"day,omitempty"`
	VenuePhotos      BoolAsAnInt             `url:"venuePhotos,omitempty"`
	LastVenue        string                  `url:"lastVenue,omitempty"`
	OpenNow          BoolAsAnInt             `url:"openNow,omitempty"`
	SortByDistance   BoolAsAnInt             `url:"sortByDistance,omitempty"`
	Price            []int                   `url:"price,omitempty"`
	Saved            BoolAsAnInt             `url:"saved,omitempty"`
	Specials         BoolAsAnInt             `url:"specials,omitempty"`
}

type VenueExploreResponse struct {
	SuggestedFilters          SuggestedFilters `json:"suggestedFilters"`
	Warning                   Warning          `json:"warning"`
	SuggestedRadius           int              `json:"suggestedRadius"`
	HeaderLocation            string           `json:"headerLocation"`
	HeaderFullLocation        string           `json:"headerFullLocation"`
	HeaderLocationGranularity string           `json:"headerLocationGranularity"`
	TotalResults              int              `json:"totalResults"`
	SuggestedBounds           SuggestedBounds  `json:"suggestedBounds"`
	Groups                    []Recommendation `json:"groups"`
}

type SuggestedFilters struct {
	Header  string   `json:"header"`
	Filters []Filter `json:"filters"`
}

type Filter struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Warning struct {
	Text string `json:"text"`
}

type SuggestedBounds struct {
	Ne LatLong `json:"ne"`
	Sw LatLong `json:"sw"`
}

type LatLong struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Recommendation struct {
	Type  string      `json:"type"`
	Name  string      `json:"name"`
	Items []Recommend `json:"items"`
}

type Recommend struct {
	Reasons    Reasons `json:"reasons"`
	Venue      Venue   `json:"venue"`
	Tips       []Tip   `json:"tips"`
	ReferralID string  `json:"referralId"`
}

// Explore returns a list of recommended venues near the current location.
// https://developer.foursquare.com/docs/venues/explore
func (s *VenueService) Explore(params *VenueExploreParams) (*VenueExploreResponse, *http.Response, error) {
	if params.LatLong == "" && params.Near == "" {
		return nil, nil, errors.New("LatLong or Near are required")
	}

	exploreResponse := new(VenueExploreResponse)
	response := new(Response)

	resp, err := s.sling.New().Get("explore").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, exploreResponse)
	}

	return exploreResponse, resp, relevantError(err, *response)
}
