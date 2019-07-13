package models

type LeagueAlias struct {
	ID       uint   `gorm:"primary_key:id"`
	LeagueID uint   `gorm:"column:league_id;index:idx_alias_league"`
	Name     string `gorm:"column:name"`
	League   League
}
