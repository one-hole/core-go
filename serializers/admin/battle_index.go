package admin

import (
	"core-go/models"

	"github.com/one-hole/gonrails/serializers"
	"github.com/w-zengtao/struct2json"
)

type TeamSerializer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (s *TeamSerializer) Serialize(v interface{}) map[string]interface{} {
	if obj, ok := v.(*models.Team); ok {
		s = &TeamSerializer{
			ID:   obj.ID,
			Name: obj.Name,
		}
		ans, _ := struct2json.Convert(s)
		return ans
	}

	return map[string]interface{}{"error": "params passed to method Serialize(v interface{}) is not a Team"}
}

type BattleIndexSerializer struct {
	ID        uint                   `json:"id"`
	LeftTeam  map[string]interface{} `json:"left_team"`
	RightTeam map[string]interface{} `json:"right_team"`
}

func (s *BattleIndexSerializer) Serialize(v interface{}) map[string]interface{} {
	if obj, ok := v.(*models.Battle); ok {
		s = &BattleIndexSerializer{
			ID:        obj.ID,
			LeftTeam:  serializers.SingleSerializer(&TeamSerializer{}, &obj.LeftTeam),
			RightTeam: serializers.SingleSerializer(&TeamSerializer{}, &obj.RightTeam),
		}

		ans, _ := struct2json.Convert(s)
		return ans
	}
	return map[string]interface{}{"error": "params passed to method Serialize(v interface{}) is not a Report"}
}
