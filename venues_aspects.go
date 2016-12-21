package foursquarego

import (
	"encoding/json"
	"net/http"
)

// PhotoGroup are the group options on VenueService.Photos
type PhotoGroup string

// Options for a PhotoGroup
const (
	PhotoGroupVenue   PhotoGroup = "venue"
	PhotoGroupCheckin PhotoGroup = "checkin"
)

// VenuePhotosParams are the paremeters for the VenueService.Photos
type VenuePhotosParams struct {
	VenueID string     `url:"-"`
	Group   PhotoGroup `url:"group,omitempty"`
	Limit   int        `url:"limit,omitempty"`
	Offset  int        `url:"offset,omitempty"`
}

type venuePhotoResp struct {
	Photos PhotoGrouping `json:"photos"`
}

// Photos gets photos for a venue
// https://developer.foursquare.com/docs/venues/photos
func (s *VenueService) Photos(params *VenuePhotosParams) (*PhotoGrouping, *http.Response, error) {
	photos := new(venuePhotoResp)
	response := new(Response)

	resp, err := s.sling.New().Get(params.VenueID+"/photos").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, photos)
	}
	return &photos.Photos, resp, relevantError(err, *response)

}

type venueEventResp struct {
	Events Events `json:"events"`
}

type Events struct {
	Count   int     `json:"count"`
	Summary string  `json:"summary"`
	Items   []Event `json:"items"`
}

type Event struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
	HereNow    HereNow    `json:"hereNow"`
	AllDay     bool       `json:"allDay"`
	Date       int64      `json:"date"`
	TimeZone   string     `json:"timeZone"`
	Stats      Stats      `json:"stats"`
	URL        string     `json:"url"`
}

// Events are music and movie events at this venue
// https://developer.foursquare.com/docs/venues/events
func (s *VenueService) Events(id string) (*Events, *http.Response, error) {
	events := new(venueEventResp)
	response := new(Response)

	resp, err := s.sling.New().Get(id+"/events").Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, events)
	}

	return &events.Events, resp, relevantError(err, *response)
}

// VenueHoursResp is the response for the venue hours endpoint
type VenueHoursResp struct {
	Hours   HoursResp `json:"hours"`
	Popular HoursResp `json:"popular"`
}

// HoursResp is an struct inside the VenueHoursResp
type HoursResp struct {
	TimeFrames []HoursTimeFrame `json:"timeframes"`
}

// HoursTimeFrame is specific to the hours endpoint
// it switches the Days from a string to an array.
// https://developer.foursquare.com/docs/responses/hours
type HoursTimeFrame struct {
	Days          []int       `json:"days"`
	IncludesToday bool        `json:"includesToday"`
	Open          []HoursOpen `json:"open"`
}

// HoursOpen contains the start time and end time when the HoursTimeFrame
// is open.
// Used only in HoursTimeFrame
type HoursOpen struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Hours Returns hours for a venue.
// https://developer.foursquare.com/docs/venues/hours
func (s *VenueService) Hours(id string) (*VenueHoursResp, *http.Response, error) {
	hours := new(VenueHoursResp)
	response := new(Response)

	resp, err := s.sling.New().Get(id+"/hours").Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, hours)
	}

	return hours, resp, relevantError(err, *response)
}

type venueLikesResp struct {
	Likes LikesResp `json:"likes"`
}

// Likesresp is the response for the venue likes endpoint
type LikesResp struct {
	Count   int    `json:"count"`
	Summary string `json:"summary"`
	Items   []User `json:"items"`
	Like    bool   `json:"like"`
}

// Likes returns friends and a total count of users who have liked this venue.
// https://developer.foursquare.com/docs/venues/likes
func (s *VenueService) Likes(id string) (*LikesResp, *http.Response, error) {
	likes := new(venueLikesResp)
	response := new(Response)

	resp, err := s.sling.New().Get(id+"/likes").Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, likes)
	}

	return &likes.Likes, resp, relevantError(err, *response)
}

type venueLinkResp struct {
	Links Links `json:"links"`
}

type Links struct {
	Count int    `json:"count"`
	Items []Link `json:"items"`
}

// Link is part of the response for the venues link endpoint
// https://developer.foursquare.com/docs/responses/link
type Link struct {
	Provider Provider `json:"provider"`
	LinkedID string   `json:"linkedId"`
	URL      string   `json:"url"`
}

type Provider struct {
	ID string `json:"id"`
}

// Links returns URLs or identifies from third parties for this venue
// https://developer.foursquare.com/docs/venues/links
func (s *VenueService) Links(id string) (*Links, *http.Response, error) {
	links := new(venueLinkResp)
	response := new(Response)

	resp, err := s.sling.New().Get(id+"/links").Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, links)
	}

	return &links.Links, resp, relevantError(err, *response)
}

// ListedGroup are the group options on VenueService.Listed
type ListedGroup string

// Options for a ListedGroup
const (
	ListedGroupOther ListedGroup = "other"
)

// VenueListedParams are the parameters for VenueService.Listed
type VenueListedParams struct {
	VenueID string      `url:"-"`
	Group   ListedGroup `url:"group,omitempty"`
	Limit   int         `url:"limit,omitempty"`
	Offset  int         `url:"offset,omitempty"`
}

type venueListedResp struct {
	Lists Listed `json:"lists"`
}

// Listed returns the lists that this venue appears on
// https://developer.foursquare.com/docs/venues/listed
func (s *VenueService) Listed(params *VenueListedParams) (*Listed, *http.Response, error) {
	lists := new(venueListedResp)
	response := new(Response)

	resp, err := s.sling.New().Get(params.VenueID+"/listed").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, lists)
	}

	return &lists.Lists, resp, relevantError(err, *response)
}

type venueNextVenuesResp struct {
	NextVenues nextVenues `json:"nextVenues"`
}

type nextVenues struct {
	Count int     `json:"count"`
	Items []Venue `json:"items"`
}

// NextVenues returns venues that are checked into after the given one
// https://developer.foursquare.com/docs/venues/nextvenues
func (s *VenueService) NextVenues(id string) ([]Venue, *http.Response, error) {
	venues := new(venueNextVenuesResp)
	response := new(Response)

	resp, err := s.sling.New().Get(id+"/nextvenues").Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venues)
	}

	return venues.NextVenues.Items, resp, relevantError(err, *response)
}

type venueMenuResp struct {
	Menu MenuResp `json:"menu"`
}

type MenuResp struct {
	Provider MenuProvider `json:"provider"`
	Menus    Menus        `json:"menus"`
}

type MenuProvider struct {
	Name             string `json:"name"`
	AttributionImage string `json:"attributionImage"`
	AttributionLink  string `json;"attributionLink"`
	AttributionText  string `json:"attributionText"`
}

type Menus struct {
	Count int        `json:"count"`
	Items []FullMenu `json:"items"`
}

type FullMenu struct {
	MenuID      string  `json:"menuId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Entries     Entries `json:"entries"`
}

type Entries struct {
	Count int     `json:"count"`
	Items []Entry `json:"items"`
}

type Entry struct {
	SectionID string     `json:"sectionId"`
	Name      string     `json:"name"`
	Entries   SubEntries `json:"entries"`
}

type SubEntries struct {
	Count int        `json:"count"`
	Items []SubEntry `json:"items"`
}

type SubEntry struct {
	EntryID     string   `json:"entryId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Prices      []string `json:"prices"`
	Price       string   `json:"price"`
	Options     Omitted  `json:"options"`
	Additions   Omitted  `json:"additions"`
}

// Menu returns menu information for a venue.
// https://developer.foursquare.com/docs/venues/menu
func (s *VenueService) Menu(id string) (*MenuResp, *http.Response, error) {
	menuResp := new(venueMenuResp)
	response := new(Response)

	resp, err := s.sling.New().Get(id+"/menu").Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, menuResp)
	}

	return &menuResp.Menu, resp, relevantError(err, *response)
}

// TipSort is the sort options on VenueService.Tips
type TipSort string

// Options for TipSort
const (
	TipsSortFriends TipSort = "friends"
	TipSortRecent   TipSort = "recent"
	TipSortPopular  TipSort = "popular"
)

// VenueTipsParams are the parameters for VenueService.Tips
type VenueTipsParams struct {
	VenueID string  `url:"-"`
	Sort    TipSort `url:"sort,omitempty"`
	Limit   int     `url:"limit,omitempty"`
	Offset  int     `url:"offset,omitempty"`
}

type tipResp struct {
	Tips tipsResp `json:"tips"`
}

type tipsResp struct {
	Count int   `json:"count"`
	Items []Tip `json:"items"`
}

// Tips returns tips for a venue.
// https://developer.foursquare.com/docs/venues/tips
func (s *VenueService) Tips(params *VenueTipsParams) ([]Tip, *http.Response, error) {
	tipResp := new(tipResp)
	response := new(Response)

	resp, err := s.sling.New().Get(params.VenueID+"/tips").QueryStruct(params).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, tipResp)
	}

	return tipResp.Tips.Items, resp, relevantError(err, *response)
}
