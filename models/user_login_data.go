package models

import (
	"time"
)

type UserLoginData struct {
	ID                      int64     `json:"user_id" gorm:"column:user_id"`
	PasswordHash            string    `json:"password_hash"`
	PasswordSalt            string    `json:"password_salt"`
	HashingAlgorithmId      int64     `json:"hashing_algorithms_id" gorm:"column:hashing_algorithms_id"`
	ConfirmationToken       string    `json:"confirmation_token"`
	TokenGenerationTime     time.Time `json:"token_generation_time"`
	PasswordRecoveryToken   string    `json:"password_recovery_token"`
	RecoveryTokenTime       time.Time `json:"recovery_token_time"`
	EmailValidationStatusId int64     `json:"email_validation_status_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`

	HashAlgorithm HashingAlgorithms `gorm:"foreignKey:hashing_algorithms_id;references:hashing_algorithms_id"`
}

// TableName gives table name of model
func (u UserLoginData) TableName() string {
	return "user_login_data"
}
