package models

import (
	
	"time"
)

type Booked_properties struct {
	Id          string `gorm:"primaryKey;type:uuid"`
	UserId      string
	PropertyId    string
	Created_at  time.Time
}
