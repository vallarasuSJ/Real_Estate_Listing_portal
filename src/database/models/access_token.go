package models

import "time"

type AccessToken struct {
	Token         string
	RefreshTokens string
	UserId     string `gorm:"type:uuid"`
	ExpiresAt     time.Time
}
