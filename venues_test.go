package foursquarego

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestVenueService_Details(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/40a55d80f964a52020f31ee3", func(w http.ResponseWriter, r *http.Request) {

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
	venue, _, err := client.Venues.Details("40a55d80f964a52020f31ee3")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(venue)
}
