package models

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

/******************************************************************************************************

说明：系列赛模型
作者：w-zengtao

belongs_to Game
official id : steam_id & hltv_id

User Story
1. 每场系列赛由一到多场（Format）比赛组成
2. 每场系列赛只有左右两支队伍进行比赛
3. Live 表示比赛是否是进行中的状态
4. 系列赛需要有大比分
5. 系列赛需要有当前正在进行中的比赛场数（game_no）（未开始 0 ， 打完 == Format + 1）
6. Status 表示系列赛的状态、取消、未开始、进行中、已结束（没有特别直观的意义）
******************************************************************************************************/

type Battle struct {
	ID          uint `gorm:"primary_key;column:id"`
	GameID      uint `gorm:"column:game_id"`
	LeftTeamID  uint `gorm:"column:left_team_id"`
	RightTeamID uint `gorm:"column:right_team_id"`
	Format      uint `gorm:"column:format;default:1"` // BO1 As Default
	OfficialID  uint `gorm:"column:official_id"`
	LeagueID    uint `gorm:"column:league_id"`
	Live        bool `gorm:"default:false"`
	StartTime   time.Time
	Status      int
	LeftScore   int
	RightScore  int
	GameNO      int

	LeftTeam  Team
	RightTeam Team
	League    League
	Matches   []Match
}

// FilterParams - for scopes
func (v *Battle) FilterParams(params map[string]string) *gorm.DB {
	var records = db.Model(v)

	if value, ok := params["game_id"]; ok {
		id, _ := strconv.Atoi(value)
		records = records.Where(&Battle{GameID: uint(id)})
	}

	if value, ok := params["league_id"]; ok {
		id, _ := strconv.Atoi(value)
		records = records.Where(&Battle{LeagueID: uint(id)})
	}

	if value, ok := params["page"]; ok {
		page, _ := strconv.Atoi(value)
		records = records.Offset((page - 1) * 10).Limit(10)
	}
	return records
}
