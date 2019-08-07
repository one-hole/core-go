package games

import (
	"core-go/controllers"
	"core-go/models"
	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	fn := func(ctx *gin.Context)  {
		games := make([]*models.Game, 0)

		if err := models.DB().Model(&models.Game{}).Find(&games).Error; err == nil {
			ctx.JSON(400, controllers.ServerError())
			return
		}

		ctx.JSON(200, gin.H{
			"data": games,
		})
	}

	return gin.HandlerFunc(fn)
}