package foursquarego_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/peppage/foursquarego"
)

var CLIENT_ID = os.Getenv("CLIENT_ID")
var CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
var OAUTH_TOKEN = os.Getenv("OAUTH_TOKEN")

var api *foursquarego.FoursquareApi

func init() {
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
}

// Test cna fetch venu & marshalling is OK
func Test_FoursquareApi_Venue(t *testing.T) {
	const venueId = "40a55d80f964a52020f31ee3"
	const venueName = "Clinton St. Baking Co. & Restaurant"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	v, err := api.GetVenue(venueId)
	if err != nil {
		t.Errorf("Getting a Venue returned error: %s", err.Error())
	}
	if v.Name != venueName {
		t.Errorf("Venue %s contained incorrect text. Recieved %s", venueId, v.Name)
	}
}

func Test_FoursquareApi_VenuePhotos(t *testing.T) {
	const venueId = "40a55d80f964a52020f31ee3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	p, err := api.GetVenuePhotos(venueId, nil)
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
	e, err := api.GetVenueEvents(venueId)
	if err != nil {
		t.Errorf("Getting venue events returned error: %s", err.Error())
	}
	if len(e) < 1 {
		t.Errorf("Exepected 1 or more event, found %d", len(e))
	}
}

func Test_FoursquareApi_Categories(t *testing.T) {
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	c, err := api.GetCategories()
	if err != nil {
		t.Errorf("Getting categories returned error: %s", err.Error())
	}
	if len(c) == 0 {
		t.Errorf("Get categories returned no categories")
	}
}

func Test_FoursquareApi_VenueHereNow(t *testing.T) {
	const venueId = "4e5c0c64183883e00c042670"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	api.SetOauthToken(OAUTH_TOKEN)
	h, err := api.GetVenueHereNow(venueId, nil)
	if err != nil {
		t.Errorf("Getting venue here now returned error %s", err.Error())
	}
	// Not sure how to test this one
	fmt.Println(h)
}

func Test_FoursquareApi_VenueHours(t *testing.T) {
	const venueId = "40a55d80f964a52020f31ee3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	h, err := api.GetVenueHours(venueId)
	if err != nil {
		t.Errorf("Getting venue hours returned error %s", err.Error())
	}
	if len(h.Hours.Timeframes) < 1 {
		t.Errorf("Get hours returned no hours")
	}
}
