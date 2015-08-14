package foursquarego

import (
	"net/url"
	"strconv"
)

func (a FoursquareApi) GetVenue(id string) (venue Venue, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "/venues/" + id, url.Values{}, &data, _GET, response_ch}
	return data.Venue, (<-response_ch).err
}

// @param venueId string
// @param group string "venue"
// @param limit int 100
// @param offset int 100
func (a FoursquareApi) GetVenuePhotos(args ...interface{}) (photos []Photo, err error) {
	var venuID string
	group := "venue"
	limit := 100
	offset := 100
	if 1 > len(args) {
		panic("Not enough parameters")
	}
	for i, p := range args {
		switch i {
		case 0: // id
			param, ok := p.(string)
			if !ok {
				panic("first param not string")
			}
			venuID = param
		case 1: // group
			param, ok := p.(string)
			if !ok {
				panic("second param not string")
			}
			group = param
		case 2: // limit
			param, ok := p.(int)
			if !ok {
				panic("third param not int")
			}
			limit = param
		case 3: // offset
			param, ok := p.(int)
			if !ok {
				panic("fourth param not int")
			}
			offset = param
		}

	}
	var v = url.Values{}
	v.Set("group", group)
	v.Set("limit", strconv.Itoa(limit))
	v.Set("offset", strconv.Itoa(offset))
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "/venues/" + venuID + "/photos", v, &data, _GET, response_ch}
	return data.Photos.Items, (<-response_ch).err
}

func (a FoursquareApi) GetCategories() (categories []Category, err error) {
	response_ch := make(chan response)
	var data foursquareResponse
	a.queryQueue <- query{API_URL + "/venues/categories", url.Values{}, &data, _GET, response_ch}
	return data.Categories, (<-response_ch).err
}
