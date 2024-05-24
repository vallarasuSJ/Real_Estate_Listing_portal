package models

import (
	
	"time"
)

type Users struct {
	Id             string `gorm:"primaryKey;type:uuid"`
	Username       string
	Email          string
	Gender         string
	Password       string
	Contact_number string
	Role_id  string
	Created_at     time.Time
}
