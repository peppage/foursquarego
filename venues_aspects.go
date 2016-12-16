package foursquarego

import (
	"encoding/json"
	"net/http"
)

// PhotoGroupParam represents the options available to group photos on photos endpoint
type PhotoGroupParam string

// Options that are valid for the group on the photos endpoint
const (
	VenuePhotoGroup    = PhotoGroupParam("venue")
	CheckingPhotoGroup = PhotoGroupParam("checkin")
)

// VenuePhotosParams are the paremeters for the VenueService.Photos
type VenuePhotosParams struct {
	VenueID string          `url:"-"`
	Group   PhotoGroupParam `url:"group,omitempty"`
	Limit   int             `url:"limit,omitempty"`
	Offset  int             `url:"offset,omitempty"`
}

type venuePhotoResp struct {
	Photos PhotoGroup `json:"photos"`
}

// Photos gets photos for a venue
// https://developer.foursquare.com/docs/venues/photos
func (s *VenueService) Photos(params *VenuePhotosParams) (*PhotoGroup, *http.Response, error) {
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

	return &likes.Likes, resp, err
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

	return &links.Links, resp, err
}
