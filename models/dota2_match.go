package models

/******************************************************************************************************

Dota2Match:

1. BeforeCreate (Type = "Dota2Match")

 ******************************************************************************************************/

type Dota2Match struct {
	Match
}

func (Dota2Match) TableName() string {
	return "matches"
}
