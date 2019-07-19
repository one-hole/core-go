package models

import "time"

/******************************************************************************************************
PlayerTeam:

Action: {
	existence: 0,
	join: 1,
	leave: 2
}
 ******************************************************************************************************/

type PlayerTeam struct {
	PlayerID uint
	TeamID   uint
	At       time.Time
	Action   int
	Player   Player
	Team     Team
}
