package foursquarego

import (
	"encoding/json"
	"net/http"

	"github.com/dghubble/sling"
)

// VenueService provies a method for accessing Foursquare venue endpoints
type VenueService struct {
	sling *sling.Sling
}

func newVenueService(sling *sling.Sling) *VenueService {
	return &VenueService{
		sling: sling.Path("venues/"),
	}
}

// Need this since the responses all have subitems.
type venueResp struct {
	Venue Venue `json:"venue"`
}

// Details gets all the data for a venue
// https://developer.foursquare.com/docs/venues/venues
func (s *VenueService) Details(id string) (*Venue, *http.Response, error) {
	response := new(Response)
	venue := new(venueResp)

	resp, err := s.sling.New().Get(id).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venue)
	}
	return &venue.Venue, resp, relevantError(err, *response)
}

// VenuePhotosParams are the paremeters for the VenueService.Photos
type VenuePhotosParams struct {
	VenueID string `url:"-"`
	Group   string `url:"group,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	Offset  int    `url:"offset,omitempty"`
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
