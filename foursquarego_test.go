package foursquarego_test

import (
	"os"
	"testing"

	"github.com/peppage/foursquarego"
)

var CLIENT_ID = os.Getenv("CLIENT_ID")
var CLIENT_SECRET = os.Getenv("CLIENT_SECRET")

var api *foursquarego.FoursquareApi

func init() {
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
}

// Test cna fetch venu & marshalling is OK
func Test_FoursquareApi_Venue(t *testing.T) {
	const venueId = "40a55d80f964a52020f31ee3"
	const venueName = "Clinton St. Baking Co. & Restaurant"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	v, err := api.Venue(venueId)
	if err != nil {
		t.Errorf("Getting a Venue returned error: %s", err.Error())
	}
	if v.Name != venueName {
		t.Errorf("Venue %s contained incorrect text. Recieved %s", venueId, v.Name)
	}
}

func Test_FoursquareApi_VenuPhotos(t *testing.T) {
	const venueId = "40a55d80f964a52020f31ee3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	p, err := api.VenuePhotos(venueId, nil)
	if err != nil {
		t.Errorf("Getting venue photos returned error: %s", err.Error())
	}
	if len(p) < 1 {
		t.Errorf("Expected 1 or more photo, found %d", len(p))
	}
}

// Events are very
func Test_FoursquareApi_VenueEvents(t *testing.T) {
	const venueId = "451d2920f964a5208d3a1fe3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	e, err := api.VenueEvents(venueId)
	if err != nil {
		t.Errorf("Getting venue events returned error: %s", err.Error())
	}
	if len(e) < 1 {
		t.Errorf("Exepected 1 or more event, found %d", len(e))
	}
}

func Test_FoursquareApi_Categories(t *testing.T) {
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	c, err := api.Categories()
	if err != nil {
		t.Errorf("Getting categories returned error: %s", err.Error())
	}
	if len(c) == 0 {
		t.Errorf("Get categories returned no categories")
	}
}
