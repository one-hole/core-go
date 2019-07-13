package models

import (
	"time"
)

/******************************************************************************************************

说明：联赛模型
作者：w-zengtao

belongs_to Game
official id : steam_id & hltv_id

User Story
联赛归属于某个游戏
联赛有许多的系列赛（赛程）
有些联赛可以手动隐藏（默认不隐藏）
联赛需要有起止的时间
联赛需要有 Logo
******************************************************************************************************/

type League struct {
	ID         uint `gorm:"primary_key;column:id"`
	GameID     uint `gorm:"column:game_id;index:idx_game_league"`
	Name       string
	Status     int
	OfficialID uint       `gorm:"column:official_id"`
	StartAt    *time.Time `gorm:"column:start_time"`
	EndAt      *time.Time `gorm:"column:end_time"`
	Hidden     bool       `gorm:"column:hidden;default:false"`
	Logo       string
	SecondLogo string

	Game    Game
	Aliases []LeagueAlias
}

func (League) TableName() string {
	return "leagues"
}
