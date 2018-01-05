package foursquarego

import (
	"encoding/json"
	"net/http"
)

type categoriesResp struct {
	Categories []Category `json:"categories"`
}

// Categories returns a hierarchical list of categories applied to venues.
// https://developer.foursquare.com/docs/api/venues/categories
func (s *VenueService) Categories() ([]Category, *http.Response, error) {
	cats := new(categoriesResp)
	response := new(Response)
	resp, err := s.sling.New().Get("categories").Receive(response, response)

	if err == nil {
		json.Unmarshal(response.Response, cats)
	}

	return cats.Categories, resp, relevantError(err, *response)
}

// SearchIntent are the intent options on VenueService.Search
type SearchIntent string

// Options for SearchIntent
const (
	IntentCheckin SearchIntent = "checkin"
	IntentBrowse  SearchIntent = "browse"
	IntentGlobal  SearchIntent = "global"
	IntentMatch   SearchIntent = "match"
)

// VenueSearchParams are the parameters for the VenueService.Search
type VenueSearchParams struct {
	LatLong          string       `url:"ll,omitempty"`
	Near             string       `url:"near,omitempty"`
	LatLongAccuracy  int          `url:"llAcc,omitempty"`
	Altitude         int          `url:"alt,omitempty"`
	AltitudeAccuracy int          `url:"altAcc,omitempty"`
	Query            string       `url:"query,omitempty"`
	Limit            int          `url:"limit,omitempty"`
	Intent           SearchIntent `url:"intent,omitempty"`
	Radius           int          `url:"radius,omitempty"`
	Sw               string       `url:"sw,omitempty"`
	Ne               string       `url:"ne,omitempty"`
	CategoryID       []string     `url:"categoryId,omitempty"`
	URL              string       `url:"url,omitempty"`
	ProviderID       string       `url:"providerId,omitempty"`
	LinkedID         int          `url:"linkedId,omitempty"`
}

type venueSearchResp struct {
	Venues []Venue `json:"venues"`
}

// Search returns a list of venues near the current location, optionally matching a search term.
// https://developer.foursquare.com/docs/api/venues/search
func (s *VenueService) Search(params *VenueSearchParams) ([]Venue, *http.Response, error) {
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
	Radius           int    `url:"radius,omitempty"`
	Sw               string `url:"sw,omitempty"`
	Ne               string `url:"ne,omitempty"`
}

// MiniVenue is a compact Venue
// https://developer.foursquare.com/docs/api/venues/suggestcompletion
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
// https://developer.foursquare.com/docs/api/venues/suggestcompletion
func (s *VenueService) SuggestCompletion(params *VenueSuggestParams) ([]MiniVenue, *http.Response, error) {
	venues := new(venueSuggestResp)
	response := new(Response)

	resp, err := s.sling.New().Get("suggestCompletion").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venues)
	}

	return venues.MiniVenues, resp, relevantError(err, *response)
}

// VenueTrendingParams are the parameters for VenueService.Trending
type VenueTrendingParams struct {
	LatLong string `url:"ll,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	Radius  int    `url:"radius,omitempty"`
}

type venueTrendingResp struct {
	Venues []Venue `json:"venues"`
}

// Trending returns a list of venues near the current location with the most people currently checked in.
// https://developer.foursquare.com/docs/api/venues/trending
func (s *VenueService) Trending(params *VenueTrendingParams) ([]Venue, *http.Response, error) {
	venues := new(venueTrendingResp)
	response := new(Response)

	resp, err := s.sling.New().Get("trending").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venues)
	}

	return venues.Venues, resp, relevantError(err, *response)
}

// ExploreSection are the section options on VenueService.Explore
type ExploreSection string

// Options for ExploreSection
const (
	SectionFood       ExploreSection = "food"
	SectionDrink      ExploreSection = "drink"
	SectionCoffee     ExploreSection = "coffee"
	SectionShops      ExploreSection = "shops"
	SectionArts       ExploreSection = "arts"
	SectionOutdoors   ExploreSection = "outdoors"
	SectionSights     ExploreSection = "sights"
	SectionTrending   ExploreSection = "trending"
	SectionSpecials   ExploreSection = "specials"
	SectionNextVenues ExploreSection = "nextVenues"
	SectionTopPicks   ExploreSection = "topPicks"
)

// Novelty are the novelty options on VenueService.Explore
type Novelty string

// Options for Novelty
const (
	NoveltyNew Novelty = "new"
	NoveltyOld Novelty = "old"
)

// FriendVisit are the friendVisit options on VenueService.Explore
type FriendVisit string

// Options for FriendVisit
const (
	FriendVisited    FriendVisit = "visited"
	FriendNotVisited FriendVisit = "notvisited"
)

// ExploreTime are the time options on VenueService.Explore
type ExploreTime string

// Options for ExploreTime
const (
	TimeAny ExploreTime = "any"
)

// VenueExploreParams are the parameters for VenueService.Explore
type VenueExploreParams struct {
	LatLong          string         `url:"ll,omitempty"`
	Near             string         `url:"near,omitempty"`
	LatLongAccuracy  int            `url:"llAcc,omitempty"`
	Altitude         int            `url:"alt,omitempty"`
	AltitudeAccuracy int            `url:"altAcc,omitempty"`
	Radius           int            `url:"radius,omitempty"`
	Section          ExploreSection `url:"section,omitempty"`
	Query            string         `url:"query,omitempty"`
	Limit            int            `url:"limit,omitempty"`
	Offset           int            `url:"offset,omitempty"`
	Novelty          Novelty        `url:"novelty,omitempty"`
	FriendVisits     FriendVisit    `url:"friendVists,omitempty"`
	Time             ExploreTime    `url:"time,omitempty"`
	Day              ExploreTime    `url:"day,omitempty"`
	VenuePhotos      BoolAsAnInt    `url:"venuePhotos,omitempty"`
	LastVenue        string         `url:"lastVenue,omitempty"`
	OpenNow          BoolAsAnInt    `url:"openNow,omitempty"`
	SortByDistance   BoolAsAnInt    `url:"sortByDistance,omitempty"`
	Price            []int          `url:"price,omitempty"`
	Saved            BoolAsAnInt    `url:"saved,omitempty"`
	Specials         BoolAsAnInt    `url:"specials,omitempty"`
}

// VenueExploreResp is the response for VenueService.Explore
// https://developer.foursquare.com/docs/api/venues/explore
type VenueExploreResp struct {
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

// SuggestedFilters are filters to show the user
type SuggestedFilters struct {
	Header  string   `json:"header"`
	Filters []Filter `json:"filters"`
}

// Filter is a Filter in SuggestedFilters
type Filter struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Warning is a text field that contains a warning message
type Warning struct {
	Text string `json:"text"`
}

// SuggestedBounds are the bounds that were used in the search
type SuggestedBounds struct {
	Ne LatLong `json:"ne"`
	Sw LatLong `json:"sw"`
}

// LatLong simple lat/long fields for SuggestedBounds
type LatLong struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Recommendation the groups field in VenueExploreResp
type Recommendation struct {
	Type  string      `json:"type"`
	Name  string      `json:"name"`
	Items []Recommend `json:"items"`
}

// Recommend is a recommendation with a Venue for VenueService.Explore
type Recommend struct {
	Reasons    Reasons `json:"reasons"`
	Venue      Venue   `json:"venue"`
	Tips       []Tip   `json:"tips"`
	ReferralID string  `json:"referralId"`
}

// Explore returns a list of recommended venues near the current location.
// https://developer.foursquare.com/docs/api/venues/explore
func (s *VenueService) Explore(params *VenueExploreParams) (*VenueExploreResp, *http.Response, error) {
	exploreResponse := new(VenueExploreResp)
	response := new(Response)

	resp, err := s.sling.New().Get("explore").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, exploreResponse)
	}

	return exploreResponse, resp, relevantError(err, *response)
}
