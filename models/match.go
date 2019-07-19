package models

import "time"

type Match struct {
	ID       uint `gorm:"primary_key;column:id"`
	BattleID uint
	StartAt  time.Time
	EndAt    time.Time
	Reverse  bool `gorm:"default:false"`
	Battle   Battle
	Type     string
}

func (Match) TableName() string {
	return "matches"
}
