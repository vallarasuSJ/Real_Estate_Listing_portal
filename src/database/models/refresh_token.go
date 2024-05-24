package models

import "time"

type RefreshToken struct {
	Token     string
	UserId string `gorm:"type:uuid"`
	ExpiresAt time.Time
}
