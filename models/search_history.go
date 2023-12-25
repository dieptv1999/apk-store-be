package models

import "time"

type SearchHistory struct {
	ID         int64     `json:"id" gorm:"column:id"`
	SearchText string    `json:"search_text"`
	CreatedAt  time.Time `json:"created_at"`
}

func (u SearchHistory) TableName() string {
	return "search_history"
}
