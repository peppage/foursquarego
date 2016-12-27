# Foursquarego [![Build Status](https://travis-ci.org/peppage/foursquarego.svg?branch=master)](https://travis-ci.org/peppage/foursquarego) [![GoDoc](https://godoc.org/github.com/peppage/foursquarego?status.svg)](https://godoc.org/github.com/peppage/foursquarego) [![Go Report Card](https://goreportcard.com/badge/github.com/peppage/foursquarego)](https://goreportcard.com/report/github.com/peppage/foursquarego)
Foursquarego is a simple Go package for accessing the[Foursquare API](https://developer.foursquare.com/docs/).

If you find any errors please create an issue, the foursquare API has different types even for things with the same name.

## Install
    go get -u github.com/peppage/foursquarego

## Usage
```go

    httpClient := http.DefaultClient
    // When creating the client you can specify either clientSecret or the accesstoken
    client := foursquarego.NewClient(httpClient, "foursquare", "clientId", "clientSecret", "")

    // Get venue details
    venue, resp, err := client.Venues.Details("57d1efb5498e018d15de8ba3")
    
    // Search Venues
    venues, resp, err := client.Venues.Search(&VenueSearchParams{
		LatLong: "40.7,-74",
		Query:   "singlecut",
	})

```

## License
[MIT License](LICENSE.md)
