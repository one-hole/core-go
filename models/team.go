package models

type Team struct {
	ID   uint `gorm:"primary_key;column:id"`
	Name string
	Abbr string
}
