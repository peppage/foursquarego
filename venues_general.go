package foursquarego

import (
	"encoding/json"
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
