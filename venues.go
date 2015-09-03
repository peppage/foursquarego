package foursquarego

import "net/url"

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
// I'm not sure if this one is setup correctly. I get the count but the items don't seem
// to be coming through
func (a FoursquareApi) GetVenueHereNow(id string, uv url.Values) (hereNow HereNow, err error) {
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

func (a FoursquareApi) GetCategories() (categories []Category, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "venues/categories", url.Values{}, &data, _GET, response_ch}
	return data.Categories, (<-response_ch).err
}
