package battles

import (
	"core-go/models"

	"github.com/gin-gonic/gin"
)

func Index() gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		battles := make([]*models.Battle, 10)

		if err := models.DB().Model(&models.Battle{}).Find(&battles).Error; err != nil {
			return
		}
	}

	return gin.HandlerFunc(fn)
}
