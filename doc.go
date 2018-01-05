/*
Package foursquarego provides a Client for the Foursquare API.

Here are some example requests

    // You will need an http client to make all the requests.
    httpClient := http.DefaultClient

    // When creating the client you can specify either clientSecret or the accesstoken
    client := foursquarego.NewClient(httpClient, "foursquare", "clientId", "clientSecret", "")

    // Venue Details
    venue, resp, err := client.Venues.Details("57d1efb5498e018d15de8ba3")

    // Search Venues
    venues, resp, err := client.Venues.Search(&VenueSearchParams{
        LatLong: "40.7,-74",
        Query:   "singlecut",
        Intent: IntentBrowse,
    })

There is a parameters struct if there is more than just 1 parameter. If there are
strict options for the parameters then there will be a struct as seen in the search above.

For Authentication the just send either the Client Secret or the users's Access Token. If you send both
to the client it will send both to foursquare. Foursquare expects that if you're making a request
for a user you will send the Access Token. More information can be found on their auth page, https://developer.foursquare.com/docs/api/configuration/authentication

*/
package foursquarego
