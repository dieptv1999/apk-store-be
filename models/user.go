package models

import (
	"time"
)

// User model
type User struct {
	ID          int64     `json:"id"`
	UserName    string    `json:"user_name"`
	UserCode    string    `json:"user_code"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Status      int       `json:"status"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	LoginData UserLoginData `gorm:"foreignKey:ID;references:ID"`
}

// TableName gives table name of model
func (u User) TableName() string {
	return "user"
}
