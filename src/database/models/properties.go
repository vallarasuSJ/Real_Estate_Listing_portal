package models

import (
	"time"
)

type Properties struct{
	Id string `gorm:"primaryKey;type:uuid"`
	Name string
	Price int
	Location string
	UserId string 
	CategoryId string
	IsApproved bool
	IsBooked bool
	Created_at time.Time 
	
}