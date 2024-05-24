package models 

type Categories struct{
	Id string `gorm:"primaryKey;type:uuid"`
	Name string
}