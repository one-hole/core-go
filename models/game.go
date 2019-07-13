package models

type Game struct {
	ID     int    `gorm:"primary_key;column:id"`
	Name   string `gorm:"column:name"`
	NameCN string `gorm:"column:name_cn"`
}

func (Game) TableName() string {
	return "games"
}
