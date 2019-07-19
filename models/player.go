package models

type Player struct {
	ID         uint
	Name       string
	Abbr       string
	Avatar     string
	TeamID     string
	OfficialId uint
	CurTeam    Team // 当前的队伍
}
