package models

import "time"

type Ads struct {
	ID       int64     `json:"id"`
	Cover    string    `json:"cover"`
	Title    string    `json:"title"`
	CreateDt time.Time `json:"create_dt"`
	Click    int64     `json:"click"`
	Link     string    `json:"link"`
	Active   string    `json:"active"`
	Count    int64     `json:"count"`
}

// TableName gives table name of model
func (u Ads) TableName() string {
	return "ads"
}
