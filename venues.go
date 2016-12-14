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
