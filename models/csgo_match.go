package models

/******************************************************************************************************
CsgoMatch:

1. BeforeCreate (Type = "CsgoMatch")

 ******************************************************************************************************/

type CsgoMatch struct {
	Match
}

func (CsgoMatch) TableName() string {
	return "matches"
}
