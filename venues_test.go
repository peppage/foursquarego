package foursquarego

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVenueService_Details(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/5414d0a6498ea3d31a3c64cf", func(w http.ResponseWriter, r *http.Request) {

		assertMethod(t, "GET", r)

		// Open file with sample json
		f, err := os.Open("./json/venues/details.json")
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient)
	venue, _, err := client.Venues.Details("5414d0a6498ea3d31a3c64cf")
	assert.Nil(t, err)

	assert.Equal(t, "5414d0a6498ea3d31a3c64cf", venue.ID)
	assert.Equal(t, "Threes Brewing", venue.Name)

	assert.Equal(t, "7185222110", venue.Contact.Phone)
	assert.Equal(t, "(718) 522-2110", venue.Contact.FormattedPhone)
	assert.Equal(t, "threesbrewing", venue.Contact.Twitter)
	assert.Equal(t, "1494258594141562", venue.Contact.Facebook)

	assert.Equal(t, "333 Douglass St", venue.Location.Address)
	assert.Equal(t, "at 4th Ave", venue.Location.CrossStreet)
	assert.Equal(t, 40.679935578556695, venue.Location.Lat)
	assert.Equal(t, -73.98211049521852, venue.Location.Lng)
	assert.Equal(t, "display", venue.Location.LabeledLatLngs[0].Label)
	assert.Equal(t, 40.679935578556695, venue.Location.LabeledLatLngs[0].Lat)
	assert.Equal(t, -73.98211049521852, venue.Location.LabeledLatLngs[0].Lng)
	assert.Equal(t, "11217", venue.Location.PostalCode)
	assert.Equal(t, "US", venue.Location.Cc)
	assert.Equal(t, "Brooklyn", venue.Location.City)
	assert.Equal(t, "NY", venue.Location.State)
	assert.Equal(t, "United States", venue.Location.Country)
	assert.Equal(t, "333 Douglass St (at 4th Ave)", venue.Location.FormattedAddress[0])
	assert.Equal(t, "Brooklyn, NY 11217", venue.Location.FormattedAddress[1])

	assert.Equal(t, "https://foursquare.com/v/threes-brewing/5414d0a6498ea3d31a3c64cf", venue.CanonicalURL)

	assert.Equal(t, "50327c8591d4c4b30a586d5d", venue.Categories[0].ID)
	assert.Equal(t, "Brewery", venue.Categories[0].Name)
	assert.Equal(t, "Breweries", venue.Categories[0].PluralName)
	assert.Equal(t, "Brewery", venue.Categories[0].ShortName)
	assert.Equal(t, "https://ss3.4sqi.net/img/categories_v2/food/brewery_", venue.Categories[0].Icon.Prefix)
	assert.Equal(t, ".png", venue.Categories[0].Icon.Suffix)
	assert.Equal(t, true, venue.Categories[0].Primary)

	assert.Equal(t, "4bf58dd8d48988d116941735", venue.Categories[1].ID)
	assert.Equal(t, "Bar", venue.Categories[1].Name)
	assert.Equal(t, "Bars", venue.Categories[1].PluralName)
	assert.Equal(t, "Bar", venue.Categories[1].ShortName)
	assert.Equal(t, "https://ss3.4sqi.net/img/categories_v2/nightlife/pub_", venue.Categories[1].Icon.Prefix)
	assert.Equal(t, ".png", venue.Categories[1].Icon.Suffix)
	assert.Equal(t, false, venue.Categories[1].Primary)

	assert.Equal(t, true, venue.Verified)
}
