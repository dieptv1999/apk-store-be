package dto

type ApkDto struct {
	ID                       int64         `json:"id"`
	Title                    string        `json:"title"`
	Description              string        `json:"description"`
	DescriptionHTML          string        `json:"descriptionHTML"`
	Summary                  string        `json:"summary"`
	Installs                 string        `json:"installs"`
	MinInstalls              int32         `json:"minInstalls"`
	RealInstalls             int32         `json:"realInstalls"`
	Score                    float64       `json:"score"`
	Ratings                  int32         `json:"ratings"`
	Reviews                  int32         `json:"reviews"`
	Histogram                []int         `json:"histogram"`
	Price                    int32         `json:"price"`
	Free                     bool          `json:"free"`
	Currency                 string        `json:"currency"`
	Sale                     bool          `json:"sale"`
	SaleTime                 string        `json:"saleTime"`
	OriginalPrice            int32         `json:"originalPrice"`
	InAppProductPrice        string        `json:"inAppProductPrice"`
	Developer                string        `json:"developer"`
	DeveloperID              string        `json:"developerId"`
	DeveloperEmail           string        `json:"developerEmail"`
	DeveloperWebsite         string        `json:"developerWebsite"`
	DeveloperAddress         string        `json:"developerAddress"`
	PrivacyPolicy            string        `json:"privacyPolicy"`
	Genre                    string        `json:"genre"`
	GenreID                  string        `json:"genreId"`
	Icon                     string        `json:"icon"`
	HeaderImage              string        `json:"headerImage"`
	Screenshots              []string      `json:"screenshots"`
	Video                    string        `json:"video"`
	VideoImage               string        `json:"videoImage"`
	ContentRating            string        `json:"contentRating"`
	ContentRatingDescription string        `json:"contentRatingDescription"`
	AdSupported              bool          `json:"adSupported"`
	ContainsAds              bool          `json:"containsAds"`
	Released                 string        `json:"released"`
	Updated                  int64         `json:"updated"`
	Version                  string        `json:"version"`
	Comments                 []string      `json:"comments"`
	AppID                    string        `json:"appId"`
	URL                      string        `json:"url"`
	Categories               []CategoryDto `json:"categories"`
}
