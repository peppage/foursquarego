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

// Details gets all the data for a venue
// https://developer.foursquare.com/docs/venues/venues
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
// https://developer.foursquare.com/docs/responses/venue
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
	//Specials      Specials   `json:"specials"` Come back to this one not enough data
	Photos      Photos      `json:"photos"`
	VenuePage   VenuePage   `json:"venuePage"`
	Reasons     Reasons     `json:"reasons"`
	Description string      `json:"description"`
	StoreID     string      `json:"storeId"`
	Page        Page        `json:"page"`
	HereNow     HereNow     `json:"hereNow"`
	CreatedAt   int64       `json:"createdAt"`
	Tips        Tips        `json:"tips"`
	Tags        []string    `json:"tags"`
	ShortURL    string      `json:"shortUrl"`
	TimeZone    string      `json:"timeZone"`
	Listed      Listed      `json:"listed"`
	Phrases     []Phrase    `json:"phrases"`
	Hours       Hours       `json:"hours"`
	Popular     Hours       `json:"popular"`
	PageUpates  PageUpdates `json:"pageUpdates"`
	Inbox       Inbox       `json:"inbox"`
	ReferralID  string      `json:"referralId"`
	//VenueChains   Omit       `json:"-"` Not enough data
	HasPerk    bool       `json:"hasPerk"`
	Attributes Attributes `json:"attributes"`
	BestPhoto  Photo      `json:"bestPhoto"`
}

type Contact struct {
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
}

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
	IsFuzzed         bool             `json:"isFuzzed"`
	Distance         int              `json:"distance"`
}

type LabeledLatLngs struct {
	Label string  `json"label"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

type Category struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	PluralName string     `json:"pluralName"`
	ShortName  string     `json:"shortName"`
	Icon       Icon       `json:"icon"`
	Primary    bool       `json:"primary"`
	Categories []Category `json:"categories,omitempty"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Stats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
	VisitsCount   int `json:"visitsCount"`
}

type Price struct {
	Tier     int    `json:"tier"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
}

type Likes struct {
	Count   int         `json:"count"`
	Groups  []LikeGroup `json:"groups"`
	Summary string      `json:"summary"`
}

type LikeGroup struct {
	Group
	Items []User `json""items"`
}

type Menu struct {
	Type      string `json:"type"`
	Label     string `json:"label"`
	Anchor    string `json:"anchor"`
	URL       string `json:"url"`
	MobileURL string `json:"mobileUrl"`
}

type FriendVisits struct {
	Count   int               `json:"count"`
	Summary string            `json:"summary"`
	Items   []FriendVisitItem `json:"items"`
}

type FriendVisitItem struct {
	VisitedCount int  `json:"visitedCount"`
	Liked        bool `json:"liked"`
	Disliked     bool `json:"disliked"`
	Oked         bool `json:"oked"`
	User         User `json:"user"`
}

type User struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Gender       string    `json:"gender"`
	Relationship string    `json:"relationship"`
	Photo        *Photo    `json:"photo"`
	Type         string    `json:"type"`
	Venue        VenuePage `json:"venue"`
	Tips         UserTips  `json:"tips"`
	Lists        Lists     `json:"lists"`
	HomeCity     string    `json:"homeCity"`
	Bio          string    `json:"bio"`
	Contact      Contact   `json:"contact"`
}

type UserTips struct {
	Count int `json:"count"`
}

type Lists struct {
	Groups []Group `json:"groups"`
}

type BeenHere struct {
	Count                int   `json:"count"`
	UnconfirmedCount     int   `json:"unconfirmedCount"`
	Marked               bool  `json:"marked"`
	LastVisitedAt        int64 `json:"lastVisitedAt"`
	LastCheckinExpiredAt int64 `json:"lastCheckinExpiredAt"`
}

type Photos struct {
	Count  int          `json:"count"`
	Groups []PhotoGroup `json:"groups"`
}

type PhotoGroup struct {
	Group
	Items []Photo `json:"items"`
}

// Photo is a foursquare photo
// https://developer.foursquare.com/docs/responses/photo.html
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

type PhotoSource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VenuePage struct {
	ID string `json:"id"`
}

type Reasons struct {
	Count int      `json:"count"`
	Items []Reason `json:"items"`
}

type Reason struct {
	Summary    string       `json:"summary"`
	Type       string       `json:"type"`
	ReasonName string       `json:"reasonName"`
	Message    string       `json:"message"`
	Target     ReasonTarget `json:"target"`
	Count      int          `json:"count"`
}

type ReasonTarget struct {
	Type   string       `json:"type"`
	Object ReasonObject `json:"object"`
}

type ReasonObject struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Target     ReasonObjectTarget `json:"target"`
	Ignoreable bool               `json:"ignoreable"`
}

type ReasonObjectTarget struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Page struct {
	User User `json:"user"`
}

type HereNow struct {
	Count   int            `json:"count"`
	Summary string         `json:"summary"`
	Groups  []HereNowGroup `json:"Groups"`
}

type HereNowGroup struct {
	Group
	//Items
}

type Tips struct {
	Count  int        `json:"count"`
	Groups []TipGroup `json:"groups"`
}

type TipGroup struct {
	Group
	Items []Tip `json:"items"`
}

type Tip struct {
	ID           string `json:"id"`
	CreatedAt    int    `json:"createdAt"`
	Text         string `json:"text"`
	Type         string `json:"type"`
	URL          string `json:"url"`
	CanonicalURL string `json:"canonicalurl"`
	Photo        Photo  `json:"photo"`
	PhotoURL     string `json:"photoUrl"`
	//Flags        Omit   `json:"flags"` //TODO:  not enough data
	Likes                 Likes  `json:"likes"`
	Like                  bool   `json:"like"`
	LogView               bool   `json:"logView"`
	ViewCount             int    `json:"viewCount"`
	Listed                Lists  `json:"listed"`
	AgreeCount            int    `json:"agreeCount"`
	DisagreeCount         int    `json:"disagreeCount"`
	Todo                  Count  `json:"todo"`
	User                  User   `json:"user"`
	AuthorInteractionType string `json:"authorInteractionType"`
}

type Listed struct {
	Count  int         `json:"count"`
	Groups []ListGroup `json:"groups"`
}

type ListGroup struct {
	Group
	Items []List `json:"items"`
}

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

type ListItems struct {
	Count int        `json:"count"`
	Items []ListItem `json:"items"`
}

type ListItem struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Tip       Tip    `json:"tip"`
	Photo     Photo  `json:"photo"`
}

type Phrase struct {
	Phrase string `json:"phrase"`
	Sample Sample `json:"sample"`
	Count  int    `json:"count"`
}

type Sample struct {
	Entities []Entitie `json:"entities"`
	Text     string    `json:"text"`
}

type Entitie struct {
	Indices []int  `json:"indices"`
	Type    string `json:"type"`
}

type Hours struct {
	Status         string      `json:"status"`
	IsOpen         bool        `json:"isOpen"`
	IsLocalHoliday bool        `json:"isLocalHoliday"`
	Timeframes     []TimeFrame `json:"timeframes"`
}

// TimeFrame shows when a venue is open.
type TimeFrame struct {
	Days          string `json:"days"`
	IncludesToday bool   `json:"includesToday"`
	Open          []Open `json:"open"`
	//Segments      Omit   `json:"-"` //TODO: Not enough data
}

type Open struct {
	RenderedTime string `json:"renderedTime"`
}

type PageUpdates struct {
	Count int `json:"count"`
	//Items not enough data
}

type Inbox struct {
	Count int `json:"count"`
	//Items not enough data
}

type Attributes struct {
	Groups []Attribute `json:"groups"`
}

type Attribute struct {
	Group
	Summary string          `json:"summary"`
	Items   []AttributeItem `json:"Items"`
}

type AttributeItem struct {
	DisplayName  string `json:"displayName"`
	DisplayValue string `json:"displayValue"`
	PriceTier    int    `json:"priceTier"`
}
