package foursquarego_test

import (
	"fmt"
	"testing"

	"github.com/peppage/foursquarego"
)

const CLIENT_ID = ""
const CLIENT_SECRET = ""

var api *foursquarego.FoursquareApi

func init() {
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
}

func Test_FoursquareApi_GetVenue(t *testing.T) {
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	v, _ := api.GetVenue("40a55d80f964a52020f31ee3")
}

func Test_FoursquareApi_GetCategories(t *testing.T) {
	api = foursquarego.NewFoursquareApi(CLIENT_ID, CLIENT_SECRET)
	_ = "breakpoint"
	c, _ := api.GetCategories()
	fmt.Println(c)
}
