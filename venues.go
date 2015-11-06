package foursquarego

import (
	"errors"
	"net/url"
)

func (a FoursquareApi) GetVenue(id string) (venue Venue, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id, url.Values{}, &data, _GET, response_ch}
	return data.Venue, (<-response_ch).err
}

// valid url.Values are: group, limit, offset
func (a FoursquareApi) GetVenuePhotos(id string, uv url.Values) (photos []Photo, err error) {
	uv = cleanValues(uv)
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/photos", uv, &data, _GET, response_ch}
	return data.Photos.Items, (<-response_ch).err
}

func (a FoursquareApi) GetVenueEvents(id string) (events []Event, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/events", url.Values{}, &data, _GET, response_ch}
	return data.Events.Items, (<-response_ch).err
}

// valid url.Values are: limit, offset
// This is reqlly a swarm endpoint
func (a FoursquareApi) GetVenueHereNow(id string, uv url.Values) (hereNow HereNow, err error) {
	if a.oauthToken == "" {
		return HereNow{}, errors.New("Requires Acting User")
	}
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/herenow", url.Values{}, &data, _GET, response_ch}
	return data.HereNow, (<-response_ch).err
}

func (a FoursquareApi) GetVenueHours(id string) (hours HoursResponse, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/hours", url.Values{}, &data, _GET, response_ch}
	return HoursResponse{data.Hours, data.Popular}, (<-response_ch).err
}

func (a FoursquareApi) GetVenueLikes(id string) (likes LikesResponse, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/likes", url.Values{}, &data, _GET, response_ch}
	return LikesResponse{data.Likes, data.Like}, (<-response_ch).err
}

func (a FoursquareApi) GetVenueLinks(id string) (links LinksResponse, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/links", url.Values{}, &data, _GET, response_ch}
	return data.Links, (<-response_ch).err
}

// valid url.Values are: group, limit, offset
func (a FoursquareApi) GetVenueListed(id string, uv url.Values) (lists Listed, err error) {
	uv = cleanValues(uv)
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/listed", uv, &data, _GET, response_ch}
	return data.Lists, (<-response_ch).err
}

func (a FoursquareApi) GetVenueMenu(id string) (menu MenuResponse, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/menu", url.Values{}, &data, _GET, response_ch}
	return data.Menu, (<-response_ch).err
}

func (a FoursquareApi) GetVenueSimilar(id string) (venues SimilarVenueResponse, err error) {
	if a.oauthToken == "" {
		return SimilarVenueResponse{}, errors.New("Requires Acting User")
	}
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/" + id + "/similar", url.Values{}, &data, _GET, response_ch}
	return data.SimilarVenues, (<-response_ch).err
}

func (a FoursquareApi) GetCategories() (categories []Category, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/categories", url.Values{}, &data, _GET, response_ch}
	return data.Categories, (<-response_ch).err
}

func (a FoursquareApi) Search(uv url.Values) (venues []Venue, err error) {
	uv = cleanValues(uv)
	if uv.Get("ll") == "" && uv.Get("near") == "" {
		return []Venue{}, errors.New("ll or near values required")
	}
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/search", uv, &data, _GET, response_ch}
	return data.Venues, (<-response_ch).err
}
