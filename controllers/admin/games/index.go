package games

import (
	"core-go/controllers/api_errors"
	"core-go/models"
	"net/http"

	"github.com/gonrails/gonrails/controllers/helper"

	"github.com/gin-gonic/gin"
)

// 这里的实现其实已经介入了之前的程序逻辑
func Index() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		games := make([]*models.Game, 0)

		if err := models.DB().Model(&models.Game{}).Find(&games).Error; err == nil {
			ctx.JSON(http.StatusBadRequest, api_errors.ServerError(helper.RequestString(ctx)))
			return
		}

		ctx.JSON(200, gin.H{
			"data": games,
		})
	}

	return gin.HandlerFunc(fn)
}
