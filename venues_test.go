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
		assertQueryNoUser(t, map[string]string{}, r)

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

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
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

	assert.Equal(t, 9988, venue.Stats.CheckinsCount)
	assert.Equal(t, 8301, venue.Stats.UsersCount)
	assert.Equal(t, 124, venue.Stats.TipCount)
	assert.Equal(t, 15919, venue.Stats.VisitsCount)

	assert.Equal(t, "http://www.threesbrewing.com", venue.URL)

	assert.Equal(t, 2, venue.Price.Tier)
	assert.Equal(t, "Moderate", venue.Price.Message)
	assert.Equal(t, "$", venue.Price.Currency)

	assert.Equal(t, true, venue.HasMenu)

	assert.Equal(t, 767, venue.Likes.Count)
	assert.Equal(t, "767 Likes", venue.Likes.Summary)

	assert.Equal(t, false, venue.Like)
	assert.Equal(t, false, venue.Dislike)
	assert.Equal(t, false, venue.Ok)
	assert.Equal(t, 9.4, venue.Rating)
	assert.Equal(t, "00B551", venue.RatingColor)
	assert.Equal(t, 943, venue.RatingSignals)

	assert.Equal(t, "Menu", venue.Menu.Type)
	assert.Equal(t, "Menu", venue.Menu.Label)
	assert.Equal(t, "View Menu", venue.Menu.Anchor)
	assert.Equal(t, "https://foursquare.com/palacediner/menu", venue.Menu.URL)
	assert.Equal(t, "https://foursquare.com/v/4ba37630f964a520103f38e3/device_menu", venue.Menu.MobileURL)

	assert.Equal(t, true, venue.AllowMenuURLEdit)

	assert.Equal(t, 4, venue.FriendVisits.Count)
	assert.Equal(t, "You and 3 friends have been here", venue.FriendVisits.Summary)
	assert.Equal(t, 2, venue.FriendVisits.Items[0].VisitedCount)
	assert.Equal(t, false, venue.FriendVisits.Items[0].Liked)
	assert.Equal(t, false, venue.FriendVisits.Items[0].Disliked)
	assert.Equal(t, false, venue.FriendVisits.Items[0].Oked)
	assert.Equal(t, "68150", venue.FriendVisits.Items[0].User.ID)
	assert.Equal(t, "Michael", venue.FriendVisits.Items[0].User.FirstName)
	assert.Equal(t, "Peppler", venue.FriendVisits.Items[0].User.LastName)
	assert.Equal(t, "male", venue.FriendVisits.Items[0].User.Gender)
	assert.Equal(t, "self", venue.FriendVisits.Items[0].User.Relationship)
	assert.Equal(t, "https://irs3.4sqi.net/img/user/", venue.FriendVisits.Items[0].User.Photo.Prefix)
	assert.Equal(t, "/68150-NB43B0NAABATDOBQ", venue.FriendVisits.Items[0].User.Photo.Suffix)

	assert.Equal(t, 2, venue.BeenHere.Count)
	assert.Equal(t, 0, venue.BeenHere.UnconfirmedCount)
	assert.Equal(t, true, venue.BeenHere.Marked)
	assert.Equal(t, int64(1444526165), venue.BeenHere.LastVisitedAt)
	assert.Equal(t, int64(1444536965), venue.BeenHere.LastCheckinExpiredAt)

	assert.Equal(t, 410, venue.Photos.Count)
	assert.Equal(t, "venue", venue.Photos.Groups[0].Type)
	assert.Equal(t, "Venue photos", venue.Photos.Groups[0].Name)
	assert.Equal(t, 410, venue.Photos.Groups[0].Count)
	assert.Equal(t, "549ecb0f11d2ed4887ba35ab", venue.Photos.Groups[0].Items[0].ID)
	assert.Equal(t, 1419692815, venue.Photos.Groups[0].Items[0].CreatedAt)
	assert.Equal(t, "Foursquare Web", venue.Photos.Groups[0].Items[0].Source.Name)
	assert.Equal(t, "https://foursquare.com", venue.Photos.Groups[0].Items[0].Source.URL)
	assert.Equal(t, "https://irs1.4sqi.net/img/general/", venue.Photos.Groups[0].Items[0].Prefix)
	assert.Equal(t, "/95760005_78vNYkB4sZbQ23LykVYIccyi2zSkD98qo3CHkQ-vI5k.jpg", venue.Photos.Groups[0].Items[0].Suffix)
	assert.Equal(t, 870, venue.Photos.Groups[0].Items[0].Width)
	assert.Equal(t, 580, venue.Photos.Groups[0].Items[0].Height)
	assert.Equal(t, false, venue.Photos.Groups[0].Items[0].Demoted)
	assert.Equal(t, "public", venue.Photos.Groups[0].Items[0].Visibility)

	assert.Equal(t, "95760005", venue.VenuePage.ID)

	assert.Equal(t, 1, venue.Reasons.Count)
	assert.Equal(t, "Valerie left a tip here", venue.Reasons.Items[0].Summary)
	assert.Equal(t, "social", venue.Reasons.Items[0].Type)
	assert.Equal(t, "friendTipReason", venue.Reasons.Items[0].ReasonName)
	assert.Equal(t, "Valerie left a tip here", venue.Reasons.Items[0].Message)
	assert.Equal(t, "navigation", venue.Reasons.Items[0].Target.Type)
	assert.Equal(t, "5850712ece66aa573a718988", venue.Reasons.Items[0].Target.Object.ID)
	assert.Equal(t, "venueTips", venue.Reasons.Items[0].Target.Object.Type)
	assert.Equal(t, "path", venue.Reasons.Items[0].Target.Object.Target.Type)
	assert.Equal(t, "/venues/5414d0a6498ea3d31a3c64cf/tips", venue.Reasons.Items[0].Target.Object.Target.URL)
	assert.Equal(t, false, venue.Reasons.Items[0].Target.Object.Ignoreable)
	assert.Equal(t, 1, venue.Reasons.Items[0].Count)

	assert.Equal(t, "Brewery, Bar & Event Space located in the Gowanus section of Brooklyn", venue.Description)
	assert.Equal(t, "", venue.StoreID)

	assert.Equal(t, "venuePage", venue.Page.User.Type)
	assert.Equal(t, "Brooklyn, NY", venue.Page.User.HomeCity)
	assert.Equal(t, 4, venue.Page.User.Tips.Count)
	assert.Equal(t, "", venue.Page.User.Bio)

	assert.Equal(t, 1, venue.HereNow.Count)
	assert.Equal(t, "One other person is here", venue.HereNow.Summary)

	assert.Equal(t, int64(1410650278), venue.CreatedAt)

	assert.Equal(t, 124, venue.Tips.Count)
	assert.Equal(t, "523e7cb811d2c579d576a7d7", venue.Tips.Groups[0].Items[0].ID)
	assert.Equal(t, 1379826872, venue.Tips.Groups[0].Items[0].CreatedAt)
	assert.Equal(t, "Fried pork sandwich is one of the best lunch options in the area. Extremely tasty and spicy if you're into that sort of thing", venue.Tips.Groups[0].Items[0].Text)
	assert.Equal(t, "user", venue.Tips.Groups[0].Items[0].Type)
	assert.Equal(t, "https://foursquare.com/item/523e7cb811d2c579d576a7d7", venue.Tips.Groups[0].Items[0].CanonicalURL)
	assert.Equal(t, "others", venue.Tips.Groups[0].Items[0].Likes.Groups[0].Type)
	assert.Equal(t, false, venue.Tips.Groups[0].Items[0].Like)
	assert.Equal(t, 619, venue.Tips.Groups[0].Items[0].ViewCount)
	assert.Equal(t, 0, venue.Tips.Groups[0].Items[0].AgreeCount)
	assert.Equal(t, 0, venue.Tips.Groups[0].Items[0].DisagreeCount)
	assert.Equal(t, "liked", venue.Tips.Groups[0].Items[0].AuthorInteractionType)

	assert.Equal(t, "beef", venue.Tags[0])

	assert.Equal(t, "http://4sq.com/1qxRLL3", venue.ShortURL)
	assert.Equal(t, "America/New_York", venue.TimeZone)

	assert.Equal(t, 659, venue.Listed.Count)
	assert.Equal(t, "561ffa6a498ed502d2869fa5", venue.Listed.Groups[0].Items[0].ID)
	assert.Equal(t, "Reasons to Love Gowanus, Brooklyn", venue.Listed.Groups[0].Items[0].Name)
	assert.Equal(t, "Fancy yourself an explorer? Gowanus, Brooklyn is an emerging neighborhood with plenty to offer. Filled with cozy restaurants, bars, & fun places to go out with friends, it's an area not to miss.", venue.Listed.Groups[0].Items[0].Description)
	assert.Equal(t, "others", venue.Listed.Groups[0].Items[0].Type)
	assert.Equal(t, false, venue.Listed.Groups[0].Items[0].Editable)
	assert.Equal(t, true, venue.Listed.Groups[0].Items[0].Public)
	assert.Equal(t, false, venue.Listed.Groups[0].Items[0].Collaborative)
	assert.Equal(t, "/foursquare/list/reasons-to-love-gowanus-brooklyn", venue.Listed.Groups[0].Items[0].URL)
	assert.Equal(t, "https://foursquare.com/foursquare/list/reasons-to-love-gowanus-brooklyn", venue.Listed.Groups[0].Items[0].CanonicalURL)
	assert.Equal(t, 1444936298, venue.Listed.Groups[0].Items[0].CreatedAt)
	assert.Equal(t, 1444941265, venue.Listed.Groups[0].Items[0].UpdatedAt)
	assert.Equal(t, true, venue.Listed.Groups[0].Items[0].LogView)
	assert.Equal(t, "neighborhood", venue.Listed.Groups[0].Items[0].GuideType)
	assert.Equal(t, true, venue.Listed.Groups[0].Items[0].Guide)
	assert.Equal(t, 68, venue.Listed.Groups[0].Items[0].Followers.Count)
	assert.Equal(t, "t5587055e498efdcb2fcce9d4", venue.Listed.Groups[0].Items[0].ListItems.Items[0].ID)

	assert.Equal(t, "rotating kitchen", venue.Phrases[0].Phrase)
	assert.Equal(t, 19, venue.Phrases[0].Sample.Entities[0].Indices[0])
	assert.Equal(t, "keyPhrase", venue.Phrases[0].Sample.Entities[0].Type)
	assert.Equal(t, 6, venue.Phrases[0].Count)

	assert.Equal(t, "Open until Midnight", venue.Hours.Status)
	assert.Equal(t, true, venue.Hours.IsOpen)
	assert.Equal(t, false, venue.Hours.IsLocalHoliday)
	assert.Equal(t, "Mon\u2013Tue", venue.Hours.Timeframes[0].Days)
	assert.Equal(t, true, venue.Hours.Timeframes[0].IncludesToday)
	assert.Equal(t, "5:00 PM\u2013Midnight", venue.Hours.Timeframes[0].Open[0].RenderedTime)

	assert.Equal(t, false, venue.Popular.IsOpen)
	assert.Equal(t, false, venue.Popular.IsLocalHoliday)
	assert.Equal(t, "Today", venue.Popular.Timeframes[0].Days)
	assert.Equal(t, true, venue.Popular.Timeframes[0].IncludesToday)
	assert.Equal(t, "6:00 PM\u201310:00 PM", venue.Popular.Timeframes[0].Open[0].RenderedTime)

	assert.Equal(t, 4, venue.PageUpates.Count)
	assert.Equal(t, 0, venue.Inbox.Count)

	assert.Equal(t, "price", venue.Attributes.Groups[0].Type)
	assert.Equal(t, "Price", venue.Attributes.Groups[0].Name)
	assert.Equal(t, "$$", venue.Attributes.Groups[0].Summary)
	assert.Equal(t, 1, venue.Attributes.Groups[0].Count)
	assert.Equal(t, "Price", venue.Attributes.Groups[0].Items[0].DisplayName)
	assert.Equal(t, "$$", venue.Attributes.Groups[0].Items[0].DisplayValue)
	assert.Equal(t, 2, venue.Attributes.Groups[0].Items[0].PriceTier)

	assert.Equal(t, "549ecb0f11d2ed4887ba35ab", venue.BestPhoto.ID)
	assert.Equal(t, 1419692815, venue.BestPhoto.CreatedAt)
	assert.Equal(t, "Foursquare Web", venue.BestPhoto.Source.Name)
	assert.Equal(t, "https://foursquare.com", venue.BestPhoto.Source.URL)
	assert.Equal(t, "https://irs1.4sqi.net/img/general/", venue.BestPhoto.Prefix)
	assert.Equal(t, "/95760005_78vNYkB4sZbQ23LykVYIccyi2zSkD98qo3CHkQ-vI5k.jpg", venue.BestPhoto.Suffix)
	assert.Equal(t, 870, venue.BestPhoto.Width)
	assert.Equal(t, 580, venue.BestPhoto.Height)
	assert.Equal(t, "public", venue.BestPhoto.Visibility)
}

func TestVenueService_Photos(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/5414d0a6498ea3d31a3c64cf/photos", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{}, r)

		// Open file with sample json
		f, err := os.Open("./json/venues/photos.json")
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

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	photos, _, err := client.Venues.Photos(&VenuePhotosParams{
		VenueID: "5414d0a6498ea3d31a3c64cf",
	})
	assert.Nil(t, err)

	assert.Equal(t, 30, photos.Count)
	assert.Equal(t, "549ecb0f11d2ed4887ba35ab", photos.Items[0].ID)
	assert.Equal(t, 1419692815, photos.Items[0].CreatedAt)
	assert.Equal(t, "Foursquare Web", photos.Items[0].Source.Name)
	assert.Equal(t, "https://foursquare.com", photos.Items[0].Source.URL)
	assert.Equal(t, 870, photos.Items[0].Width)
	assert.Equal(t, 580, photos.Items[0].Height)
	assert.Equal(t, false, photos.Items[0].Demoted)
	assert.Equal(t, "95760005", photos.Items[0].User.ID)
	assert.Equal(t, "Threes Brewing", photos.Items[0].User.FirstName)
	assert.Equal(t, "none", photos.Items[0].User.Gender)
	assert.Equal(t, "https://irs0.4sqi.net/img/user/", photos.Items[0].User.Photo.Prefix)
	assert.Equal(t, "/95760005-K35NSGGG10EE5XU2.png", photos.Items[0].User.Photo.Suffix)
	assert.Equal(t, "venuePage", photos.Items[0].User.Type)
	assert.Equal(t, "5414d0a6498ea3d31a3c64cf", photos.Items[0].User.Venue.ID)
	assert.Equal(t, "public", photos.Items[0].Visibility)
}

func TestVenueService_Events(t *testing.T) {
	const filePath = "./json/venues/events.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/40afe980f964a5203bf31ee3/events", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	events, _, err := client.Venues.Events("40afe980f964a5203bf31ee3")
	assert.Nil(t, err)

	assert.Equal(t, 26, events.Count)
	assert.Equal(t, "26 movies", events.Summary)
	assert.Equal(t, "580850f7d67c37ceeeaae676", events.Items[0].ID)
	assert.Equal(t, "Moonlight", events.Items[0].Name)
	assert.Equal(t, true, events.Items[0].AllDay)
	assert.Equal(t, int64(1477540800), events.Items[0].Date)
	assert.Equal(t, "America/New_York", events.Items[0].TimeZone)
	assert.Equal(t, 81, events.Items[0].Stats.CheckinsCount)
	assert.Equal(t, 78, events.Items[0].Stats.UsersCount)
	assert.Equal(t, "https://foursquare.com/events/movies?theater=AAORE&movie=194816&wired=true", events.Items[0].URL)
}

func TestVenueService_Hours(t *testing.T) {
	const filePath = "./json/venues/hours.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/40a55d80f964a52020f31ee3/hours", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	hours, _, err := client.Venues.Hours("40a55d80f964a52020f31ee3")
	assert.Nil(t, err)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, hours.Hours.TimeFrames[0].Days)
	assert.Equal(t, true, hours.Hours.TimeFrames[0].IncludesToday)
	assert.Equal(t, "0800", hours.Hours.TimeFrames[0].Open[0].Start)
	assert.Equal(t, "1600", hours.Hours.TimeFrames[0].Open[0].End)
}

func TestVenueService_Likes(t *testing.T) {
	const filePath = "./json/venues/likes.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/40a55d80f964a52020f31ee3/likes", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	likes, _, err := client.Venues.Likes("40a55d80f964a52020f31ee3")
	assert.Nil(t, err)

	assert.Equal(t, 1261, likes.Count)
	assert.Equal(t, "1261 Likes", likes.Summary)
	assert.Equal(t, "203153", likes.Items[0].ID)
	assert.Equal(t, "Emiliano", likes.Items[0].FirstName)
	assert.Equal(t, "Viscarra", likes.Items[0].LastName)
	assert.Equal(t, "male", likes.Items[0].Gender)
	assert.Equal(t, "https://irs1.4sqi.net/img/user/", likes.Items[0].Photo.Prefix)
	assert.Equal(t, "/203153-0BI5LE1Y2ITI4XUU.jpg", likes.Items[0].Photo.Suffix)
	assert.Equal(t, false, likes.Like)
}

func TestVenueservice_Links(t *testing.T) {
	const filePath = "./json/venues/links.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/3fd66200f964a52074e31ee3/links", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	links, _, err := client.Venues.Links("3fd66200f964a52074e31ee3")
	assert.Nil(t, err)

	assert.Equal(t, 11, links.Count)
	assert.Equal(t, "nyt", links.Items[0].Provider.ID)
	assert.Equal(t, "1002207971611", links.Items[0].LinkedID)
	assert.Equal(t, "http://www.nytimes.com/restaurants/1002207971611/db-bistro-moderne/details.html", links.Items[0].URL)
}

func TestVenueService_Categories(t *testing.T) {
	const filePath = "./json/venues/categories.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/categories", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	categories, _, err := client.Venues.Categories()
	assert.Nil(t, err)

	assert.Equal(t, "4d4b7104d754a06370d81259", categories[0].ID)
	assert.Equal(t, "Arts & Entertainment", categories[0].Name)
	assert.Equal(t, "Arts & Entertainment", categories[0].PluralName)
	assert.Equal(t, "Arts & Entertainment", categories[0].ShortName)
	assert.Equal(t, "https://ss3.4sqi.net/img/categories_v2/arts_entertainment/default_", categories[0].Icon.Prefix)
	assert.Equal(t, ".png", categories[0].Icon.Suffix)
	assert.Equal(t, "56aa371be4b08b9a8d5734db", categories[0].Categories[0].ID)
	assert.Equal(t, "Amphitheater", categories[0].Categories[0].Name)
	assert.Equal(t, "Amphitheaters", categories[0].Categories[0].PluralName)
	assert.Equal(t, "Amphitheater", categories[0].Categories[0].ShortName)
	assert.Equal(t, "https://ss3.4sqi.net/img/categories_v2/arts_entertainment/default_", categories[0].Categories[0].Icon.Prefix)
	assert.Equal(t, ".png", categories[0].Categories[0].Icon.Suffix)
}

func TestVenueService_Search(t *testing.T) {
	const filePath = "./json/venues/search.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/search", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{
			"ll":    "40.7,-74",
			"query": "singlecut",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	venues, _, err := client.Venues.Search(&VenueSearchParams{
		LatLong: "40.7,-74",
		Query:   "singlecut",
	})
	assert.Nil(t, err)

	assert.Equal(t, "4f68de6bd5fbee32e5f4f3a5", venues[0].ID)
	assert.Equal(t, false, venues[0].HasPerk)

}

func TestVenueService_Listed(t *testing.T) {
	const filePath = "./json/venues/listed.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/4f68de6bd5fbee32e5f4f3a5/listed", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{
			"limit": "1",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	lists, _, err := client.Venues.Listed(&VenueListedParams{
		VenueID: "4f68de6bd5fbee32e5f4f3a5",
		Limit:   1,
	})
	assert.Nil(t, err)

	assert.Equal(t, 382, lists.Count)
	assert.Equal(t, "others", lists.Groups[0].Type)
	assert.Equal(t, "Lists from other people", lists.Groups[0].Name)
	assert.Equal(t, 382, lists.Groups[0].Count)
	assert.Equal(t, "561e72cf498ee3be0c697a9a", lists.Groups[0].Items[0].ID)
	assert.Equal(t, "America's Best Breweries", lists.Groups[0].Items[0].Name)
	assert.Equal(t, "others", lists.Groups[0].Items[0].Type)
	assert.Equal(t, false, lists.Groups[0].Items[0].Editable)
	assert.Equal(t, true, lists.Groups[0].Items[0].Public)
	assert.Equal(t, false, lists.Groups[0].Items[0].Collaborative)
	assert.Equal(t, "/foursquare/list/americas-best-breweries", lists.Groups[0].Items[0].URL)
}

func TestVenueService_SuggestCompletion(t *testing.T) {
	const filePath = "./json/venues/suggest.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/suggestCompletion", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{
			"ll":    "40.7,-74",
			"query": "foursqu",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	venues, _, err := client.Venues.SuggestCompletion(&VenueSuggestParams{
		LatLong: "40.7,-74",
		Query:   "foursqu",
	})
	assert.Nil(t, err)

	assert.Equal(t, "4ef0e7cf7beb5932d5bdeb4e", venues[0].ID)
	assert.Equal(t, "Foursquare HQ", venues[0].Name)
	assert.Equal(t, false, venues[0].HasPerk)
}

func TestVenueService_Trending(t *testing.T) {
	const filePath = "./json/venues/trending.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/trending", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQueryNoUser(t, map[string]string{
			"ll":    "40.7,-74",
			"limit": "2",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	venues, _, err := client.Venues.Trending(&VenueTrendingParams{
		LatLong: "40.7,-74",
		Limit:   2,
	})
	assert.Nil(t, err)

	assert.Equal(t, "4eb90d85722e09311d356915", venues[0].ID)
	assert.Equal(t, "World Trade Center Transportation Hub (The Oculus)", venues[0].Name)
	assert.Equal(t, "pathtrain", venues[0].Contact.Twitter)
}

func TestVenueService_NextVenues(t *testing.T) {
	const filePath = "./json/venues/nextvenues.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/v2/venues/40a55d80f964a52020f31ee3/nextvenues", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := newClient(httpClient, "foursquare", clientID, clientSecret, "")
	venues, _, err := client.Venues.NextVenues("40a55d80f964a52020f31ee3")
	assert.Nil(t, err)

	assert.Len(t, venues, 5)
	assert.Equal(t, "4acbe67af964a52044c820e3", venues[0].ID)
}
