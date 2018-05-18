package foursquarego

import (
	"encoding/json"
	"net/http"

	"github.com/dghubble/sling"
)

// VenueService provies a method for accessing Foursquare venue endpoints
type VenueService struct {
	sling *sling.Sling
}

func newVenueService(sling *sling.Sling) *VenueService {
	return &VenueService{
		sling: sling.Path("venues/"),
	}
}

// Need special responses for endpoints since
// the objects we want are subitems of the response
type venueResp struct {
	Venue Venue `json:"venue"`
}

// SetHeader sets a header to be sent with the request for internationalization
// https://developer.foursquare.com/docs/api/configuration/versioning
func (s *VenueService) SetHeader(key, value string) *VenueService {
	s.sling.Set(key, value)
	return s
}

// Details gets all the data for a venue
// https://developer.foursquare.com/docs/api/venues/details
func (s *VenueService) Details(id string) (*Venue, *http.Response, error) {
	response := new(Response)
	venue := new(venueResp)

	resp, err := s.sling.New().Get(id).Receive(response, response)
	if err == nil {
		json.Unmarshal(response.Response, venue)
	}
	return &venue.Venue, resp, relevantError(err, *response)
}

// Venue represents a foursquare Venue.
// https://developer.foursquare.com/docs/api/venues/details
type Venue struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Contact          Contact      `json:"contact"`
	Location         Location     `json:"location"`
	CanonicalURL     string       `json:"canonicalUrl"`
	Categories       []Category   `json:"categories"`
	Verified         bool         `json:"verified"`
	Stats            Stats        `json:"stats"`
	URL              string       `json:"url"`
	Price            Price        `json:"price"`
	HasMenu          bool         `json:"hasMenu"`
	Likes            Likes        `json:"likes"`
	Like             bool         `json:"like"`
	Dislike          bool         `json:"dislike"`
	Ok               bool         `json:"ok"`
	Rating           float64      `json:"rating"`
	RatingColor      string       `json:"ratingColor"`
	RatingSignals    int          `json:"ratingSignals"`
	Menu             Menu         `json:"menu"`
	AllowMenuURLEdit bool         `json:"allowMenuUrlEdit"`
	FriendVisits     FriendVisits `json:"friendVisits"`
	BeenHere         BeenHere     `json:"beenHere"`
	Specials         Omitted      `json:"Specials"`
	Photos           Photos       `json:"photos"`
	VenuePage        ID           `json:"venuePage"`
	Reasons          Reasons      `json:"reasons"`
	Description      string       `json:"description"`
	StoreID          string       `json:"storeId"`
	Page             Page         `json:"page"`
	HereNow          HereNow      `json:"hereNow"`
	CreatedAt        int64        `json:"createdAt"`
	Tips             Tips         `json:"tips"`
	ShortURL         string       `json:"shortUrl"`
	TimeZone         string       `json:"timeZone"`
	Listed           Listed       `json:"listed"`
	Phrases          []Phrase     `json:"phrases"`
	Hours            Hours        `json:"hours"`
	Popular          Hours        `json:"popular"`
	PageUpates       PageUpdates  `json:"pageUpdates"`
	Inbox            Inbox        `json:"inbox"`
	ReferralID       string       `json:"referralId"`
	VenueChains      Omitted      `json:"venueChains"`
	HasPerk          bool         `json:"hasPerk"`
	Attributes       Attributes   `json:"attributes"`
	BestPhoto        Photo        `json:"bestPhoto"`
	Colors           Colors       `json:"colors"`
}

// Contact are details to contact this venue. Can contain all or none.
type Contact struct {
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
	Twitter        string `json:"twitter"`
	Instagram      string `json:"instagram"`
	Facebook       string `json:"facebook"`
}

// Location is a location for the venue. Can contain all or none and
// some venues have a hidden location.
type Location struct {
	Address          string           `json:"address"`
	CrossStreet      string           `json:"crossStreet"`
	Lat              float64          `json:"lat"`
	Lng              float64          `json:"lng"`
	LabeledLatLngs   []LabeledLatLngs `json:"labeledLatLngs"`
	PostalCode       string           `json:"postalCode"`
	Cc               string           `json:"cc"`
	City             string           `json:"city"`
	State            string           `json:"state"`
	Country          string           `json:"country"`
	FormattedAddress []string         `json:"formattedAddress"`
	IsFuzzed         bool             `json:"isFuzzed,omitempty"`
	Distance         int              `json:"distance,omitempty"`
}

// LabeledLatLngs is further details in the Location of a venue.
type LabeledLatLngs struct {
	Label string  `json:"label"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

// Category is a category applied to a venue
// https://developer.foursquare.com/docs/api/venues/categories
type Category struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	PluralName string     `json:"pluralName"`
	ShortName  string     `json:"shortName"`
	Icon       Icon       `json:"icon"`
	Primary    bool       `json:"primary"`
	Categories []Category `json:"categories,omitempty"`
}

// Icon is the pieces needed to construct icons at various sizes.
type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

// Stats are the stats for a venue.
type Stats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
	VisitsCount   int `json:"visitsCount"`
}

// Price is the price tier of a venue from 1 (least pricey) - 4 (most pricey).
type Price struct {
	Tier     int    `json:"tier"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
}

// Likes is a count of users who liked the venue and groups containing
// users who liked it (friends and others).
type Likes struct {
	Count   int         `json:"count"`
	Groups  []LikeGroup `json:"groups"`
	Summary string      `json:"summary"`
}

// LikeGroup is a group of users for the Likes struct.
type LikeGroup struct {
	Group
	Items []User `json:"items"`
}

// Menu contains how to access the menu for the venue.
type Menu struct {
	Type      string `json:"type"`
	Label     string `json:"label"`
	Anchor    string `json:"anchor"`
	URL       string `json:"url"`
	MobileURL string `json:"mobileUrl"`
}

// FriendVisits contains if an authed user's friends visited, includes
// self visits.
type FriendVisits struct {
	Count   int               `json:"count"`
	Summary string            `json:"summary"`
	Items   []FriendVisitItem `json:"items"`
}

// FriendVisitItem contains the user that visited along with information
// about their interaction with the venue.
type FriendVisitItem struct {
	VisitedCount int  `json:"visitedCount"`
	Liked        bool `json:"liked"`
	Disliked     bool `json:"disliked"`
	Oked         bool `json:"oked"`
	User         User `json:"user"`
}

// User is a foursquare user
// https://developer.foursquare.com/docs/api/users/details
type User struct {
	ID           string  `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Gender       string  `json:"gender"`
	Relationship string  `json:"relationship"`
	Photo        *Photo  `json:"photo"`
	Type         string  `json:"type"`
	Venue        ID      `json:"venue"`
	Tips         Count   `json:"tips"`
	Lists        Lists   `json:"lists"`
	HomeCity     string  `json:"homeCity"`
	Bio          string  `json:"bio"`
	Contact      Contact `json:"contact"`
}

// Lists are Lists on User.
type Lists struct {
	Groups []Group `json:"groups"`
}

// BeenHere contains the number of times the acting user has
// been to the venue. Absent if there is no acting user.
type BeenHere struct {
	Count                int   `json:"count"`
	UnconfirmedCount     int   `json:"unconfirmedCount"`
	Marked               bool  `json:"marked"`
	LastVisitedAt        int64 `json:"lastVisitedAt"`
	LastCheckinExpiredAt int64 `json:"lastCheckinExpiredAt"`
}

// Photos contains a count and groups of photos for the venue.
type Photos struct {
	Count  int             `json:"count"`
	Groups []PhotoGrouping `json:"groups"`
}

// PhotoGrouping is a default group with items of type photo.
type PhotoGrouping struct {
	Group
	Items []Photo `json:"items"`
}

// Photo is a foursquare photo
// https://developer.foursquare.com/docs/api/photos/details
type Photo struct {
	ID         string      `json:"id"`
	CreatedAt  int         `json:"createdAt"`
	Source     PhotoSource `json:"source"`
	Prefix     string      `json:"prefix"`
	Suffix     string      `json:"suffix"`
	Demoted    bool        `json:"demoted"`
	Width      int         `json:"width"`
	Height     int         `json:"height"`
	User       User        `json:"user"`
	Visibility string      `json:"visibility"`
}

// PhotoSource is the source on a photo struct.
type PhotoSource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Reasons why the venue is shown.
type Reasons struct {
	Count int      `json:"count"`
	Items []Reason `json:"items"`
}

// Reason is the items in Reasons.
type Reason struct {
	Summary    string       `json:"summary"`
	Type       string       `json:"type"`
	ReasonName string       `json:"reasonName"`
	Message    string       `json:"message"`
	Target     ReasonTarget `json:"target"`
	Count      int          `json:"count"`
}

// ReasonTarget where the reason would be shwon for Reason.
type ReasonTarget struct {
	Type   string       `json:"type"`
	Object ReasonObject `json:"object"`
}

// ReasonObject is Object in ReasonTarget.
type ReasonObject struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Target     ReasonObjectTarget `json:"target"`
	Ignoreable bool               `json:"ignoreable"`
}

// ReasonObjectTarget what type of target and the url for a ReasonObject.
type ReasonObjectTarget struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// Page (optional) contains the user that is the branded page
// associated with the venue.
type Page struct {
	User User `json:"user"`
}

// HereNow contains which users are here now.
type HereNow struct {
	Count   int            `json:"count"`
	Summary string         `json:"summary"`
	Groups  []HereNowGroup `json:"Groups"`
}

// HereNowGroup is the groups item in HereNow.
type HereNowGroup struct {
	Group
	Items Omitted `json:"items"`
}

// Tips contains a count and groups of tips.
type Tips struct {
	Count  int        `json:"count"`
	Groups []TipGroup `json:"groups"`
}

// TipGroup is the standard group field where the items are tips.
type TipGroup struct {
	Group
	Items []Tip `json:"items"`
}

// Tip is a foursquare tip on a venue.
type Tip struct {
	ID                    string  `json:"id"`
	CreatedAt             int     `json:"createdAt"`
	Text                  string  `json:"text"`
	Type                  string  `json:"type"`
	URL                   string  `json:"url"`
	CanonicalURL          string  `json:"canonicalurl"`
	Photo                 Photo   `json:"photo"`
	PhotoURL              string  `json:"photoUrl"`
	Flags                 Omitted `json:"flags"`
	Likes                 Likes   `json:"likes"`
	Like                  bool    `json:"like"`
	LogView               bool    `json:"logView"`
	Listed                Lists   `json:"listed"`
	AgreeCount            int     `json:"agreeCount"`
	DisagreeCount         int     `json:"disagreeCount"`
	Todo                  Count   `json:"todo"`
	User                  User    `json:"user"`
	AuthorInteractionType string  `json:"authorInteractionType"`
}

// Listed contains a count and the grouped lists
type Listed struct {
	Count  int         `json:"count"`
	Groups []ListGroup `json:"groups"`
}

// ListGroup is the standard group field where the items are of
// type List.
type ListGroup struct {
	Group
	Items []List `json:"items"`
}

// List is a foursquare list.
// https://developer.foursquare.com/docs/api/lists/details
type List struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Type          string    `json:"type"`
	User          User      `json:"user"`
	Editable      bool      `json:"editable"`
	Public        bool      `json:"public"`
	Collaborative bool      `json:"collaborative"`
	URL           string    `json:"url"`
	CanonicalURL  string    `json:"canonicalUrl"`
	CreatedAt     int       `json:"createdAt"`
	UpdatedAt     int       `json:"updatedAt"`
	Photo         Photo     `json:"photo"`
	LogView       bool      `json:"logView"`
	GuideType     string    `json:"guideType"`
	Guide         bool      `json:"guide"`
	Followers     Count     `json:"followers"`
	ListItems     ListItems `json:"listItems"`
}

// ListItems contains a count and an array of ListItem.
type ListItems struct {
	Count int        `json:"count"`
	Items []ListItem `json:"items"`
}

// ListItem contains more information about a list.
// https://developer.foursquare.com/docs/api/lists/details
type ListItem struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Tip       Tip    `json:"tip"`
	Photo     Photo  `json:"photo"`
}

// Phrase contains a phrase commonly seen with a venue's tips.
type Phrase struct {
	Phrase string `json:"phrase"`
	Sample Sample `json:"sample"`
	Count  int    `json:"count"`
}

// Sample contains an example of a Phrase being used in a tip.
type Sample struct {
	Entities []Entitie `json:"entities"`
	Text     string    `json:"text"`
}

// Entitie contains where the Phrase is in the Sample.
type Entitie struct {
	Indices []int  `json:"indices"`
	Type    string `json:"type"`
}

// Hours contains hours during the week when a venue is open.
// Used in Venue struct.
type Hours struct {
	Status         string      `json:"status"`
	IsOpen         bool        `json:"isOpen"`
	IsLocalHoliday bool        `json:"isLocalHoliday"`
	Timeframes     []TimeFrame `json:"timeframes"`
}

// TimeFrame shows when a venue is open.
type TimeFrame struct {
	Days          string  `json:"days"`
	IncludesToday bool    `json:"includesToday"`
	Open          []Open  `json:"open"`
	Segments      Omitted `json:"Segments"`
}

// Open contains how a timeframe would be written out.
// Used only in TimeFrame.
type Open struct {
	RenderedTime string `json:"renderedTime"`
}

// PageUpdates is on a Venue.
type PageUpdates struct {
	Count int     `json:"count"`
	Items Omitted `json:"items"`
}

// Inbox is on a Venue.
type Inbox struct {
	Count int     `json:"count"`
	Items Omitted `json:"items"`
}

// Attributes contains Attribute associated with a venue.
type Attributes struct {
	Groups []Attribute `json:"groups"`
}

// Attribute contains info about Venue such as price tier,
// reservations and parking.
type Attribute struct {
	Group
	Summary string          `json:"summary"`
	Items   []AttributeItem `json:"Items"`
}

// AttributeItem is the actual value shown in an Attribute.
type AttributeItem struct {
	DisplayName  string `json:"displayName"`
	DisplayValue string `json:"displayValue"`
	PriceTier    int    `json:"priceTier"`
}

// Colors is undocumented
type Colors struct {
	HighlightedColor     Color `json:"highlightColor"`
	HighlightedTextColor Color `json:"highlightTextColor"`
	AlgoVersion          int   `json:"algoVersion"`
}

// Color is undocumented and part of colors struct
type Color struct {
	PhotoID string `json:"photoId"`
	Value   int    `json:"value"`
}
