package battles

import (
	"core-go/models"
	"core-go/serializers/admin"

	"github.com/gin-gonic/gin"
	"github.com/one-hole/gonrails/controllers/helper"
	"github.com/one-hole/gonrails/serializers"
)

// Index - GET admin/index
func Index(ctx *gin.Context) {
	var battles []*models.Battle
	var battle = &models.Battle{}

	helper.Page(battle.FilterParams(helper.Params(ctx)).Preload("LeftTeam").Preload("RightTeam").Preload("League"), helper.CurrentPage(ctx)).Find(&battles)
	resp, err := serializers.CollectionSerializer(&admin.BattleIndexSerializer{}, battles)

	if err != nil {
		ctx.JSON(400, gin.H{
			"data": err,
		})
	}

	ctx.JSON(200, gin.H{
		"data": resp,
	})
}

/*
过滤参数如下:
page
game_id
league_id
begin_at
end_at
visible
*/
