package games

import (
	"log"

	"core-go/models"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	games := make([]*models.Game, 0)

	if err := models.DB().Model(&models.Game{}).Find(&games).Error; err != nil {
		log.Println("Found Error")
	}

	ctx.JSON(200, gin.H{
		"data": games,
	})
}
