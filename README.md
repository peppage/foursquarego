Foursquarego
====================

Foursquarego is a simple Go package for accessing the Foursquare API.

If you find any errors please create an issue, the foursquare API has different types even for things with the same name. I tested as much as I could and right now this fits my needs.

This code is largely based off of [Anaconda](https://github.com/ChimeraCoder/anaconda) and I have included their LICENSE in with mine.

Examples
-------------

### Authentication

Creating a client is simple:

````go
api := foursquarego.NewFoursquareApi("your client id", "your client secret")
````

If you have a user oauth token you can assign that to the api too:

````go
api.SetOauthToken("your oauth token")
````

### Queries

Queries are conducted using a pointer to an authenticated `FoursquareApi` struct.

````go
venue, _ := api.GetVenue(venueID)
fmt.Println(venue.Name)
````
Certain endpoints allow separate optional parameter; if desired, these can be passed as the final parameter. 

````go
//Perhaps we want 200 values instead of the default
uv := url.Values{}
uv.Set("limit", "200")
p, err := api.GetVenuePhotos(venueID, uv)
````

(Remember that `url.Values` is equivalent to a `map[string][]string`, if you find that more convenient notation when specifying values). Otherwise, `nil` suffices.



Error Handling
---------------------------------
Errors are returned as an `ApiError`, which satisfies the `error` interface and can be treated as a vanilla `error`. However, it also contains the additional information that may be useful in deciding how to proceed after encountering an error.


License
-----------
Foursquarego is free software licensed under the MIT license. Details provided in the LICENSE file.
