package foursquarego

import "net/url"

func (a FoursquareApi) Venue(id string) (venue Venue, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "/venues/" + id, url.Values{}, &data, _GET, response_ch}
	return data.Venue, (<-response_ch).err
}

// valid url.Values are: group, limit, offset
func (a FoursquareApi) VenuePhotos(id string, uv url.Values) (photos []Photo, err error) {
	uv = cleanValues(uv)
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "/venues/" + id + "/photos", uv, &data, _GET, response_ch}
	return data.Photos.Items, (<-response_ch).err
}

func (a FoursquareApi) VenueEvents(id string) (events []Event, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "/venues/" + id + "/events", url.Values{}, &data, _GET, response_ch}
	return data.Events.Items, (<-response_ch).err
}

func (a FoursquareApi) Categories() (categories []Category, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "/venues/categories", url.Values{}, &data, _GET, response_ch}
	return data.Categories, (<-response_ch).err
}
