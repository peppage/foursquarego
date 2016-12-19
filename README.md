# Foursquarego
Foursquarego is a simple Go package for accessing the[Foursquare API](https://developer.foursquare.com/docs/).

If you find any errors please create an issue, the foursquare API has different types even for things with the same name.

## Install
    go get -u github.com/peppage/foursquarego

## Usage
```go

    httpClient := http.DefaultClient
    client := foursquarego.NewClient("foursquare", "clientId", "clientSecret")

    // Get venue details
    venue, resp, err := client.Venues.Details("57d1efb5498e018d15de8ba3)
    
    // Search Venues
    venues, resp, err := client.Venues.Search(&VenueSearchParams{
		LatLong: "40.7,-74",
		Query:   "singlecut",
	})

```

## License
[MIT License](LICENSE.md)