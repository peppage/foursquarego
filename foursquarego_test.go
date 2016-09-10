package foursquarego_test

import (
	"net/url"
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
	uv := url.Values{}
	uv.Set("oauth_token", OAUTH_TOKEN)
	_, err := api.GetVenueHereNow(venueId, v)
	if err != nil {
		t.Errorf("Getting venue here now returned error %s", err.Error())
	}
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

func Test_FoursquareApi_VenueLikes(t *testing.T) {
	const venueId = "40a55d80f964a52020f31ee3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	l, err := api.GetVenueLikes(venueId)
	if err != nil {
		t.Errorf("Getting venue likes returned error %s", err.Error())
	}
	if len(l.Likes.Items) < 1 {
		t.Errorf("Get Likes returned no likes")
	}
}

func Test_FoursquareApi_VenueLinks(t *testing.T) {
	const venueId = "3fd66200f964a52074e31ee3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	l, err := api.GetVenueLinks(venueId)
	if err != nil {
		t.Errorf("Getting venue links returned error %s", err.Error())
	}
	if len(l.Items) < 1 {
		t.Errorf("Get links returned no links")
	}
}

func Test_FoursquareApi_VenueListed(t *testing.T) {
	const venueId = "4989af90f964a5207f521fe3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	l, err := api.GetVenueListed(venueId, nil)
	if err != nil {
		t.Errorf("Getting venue listed returned error %s", err.Error())
	}
	if len(l.Groups[0].Items) < 1 {
		t.Errorf("Get listed returned no lists")
	}
}

func Test_FoursquareApi_VenueMenu(t *testing.T) {
	const venueId = "47a1bddbf964a5207a4d1fe3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	m, err := api.GetVenueMenu(venueId)
	if err != nil {
		t.Errorf("Getting venue menu returned error %s", err.Error())
	}
	if len(m.Menus.Items[0].Entries.Items) < 1 {
		t.Errorf("Get menu returned no menus")
	}
}

func Test_FoursquareApi_VenueSimilar(t *testing.T) {
	const venueId = "40a55d80f964a52020f31ee3"
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	uv := url.Values{}
	uv.Set("oauth_token", OAUTH_TOKEN)
	s, err := api.GetVenueSimilar(venueId, v)
	if err != nil {
		t.Errorf("Getting similar venues returned error %s", err.Error())
	}
	if len(s.Items) < 1 {
		t.Errorf("Get similar venues returned no venues")
	}
}

func Test_FoursquareApi_VenueSearch(t *testing.T) {
	uv := url.Values{}
	uv.Set("ll", "40.7,-74.0")
	uv.Set("categoryId", "50327c8591d4c4b30a586d5d")
	uv.Set("intent", "browse")
	uv.Set("radius", "100000")
	venues, err := api.Search(uv)
	if err != nil {
		t.Errorf("Searching venues returned error %s", err.Error())
	}
	if len(venues) < 1 {
		t.Errorf("Searching venues returned no venues")
	}
}
