package dto

type FilterApkDto struct {
	Type        string `json:"type"`
	GenreId     string `json:"genreId"`
	AppId       string `json:"appId"`
	DeveloperId string `json:"developerId"`
	Developer   string `json:"developer"`
}
