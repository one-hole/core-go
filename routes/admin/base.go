package admin

import (
	"core-go/controllers/admin/games"
	"core-go/controllers/home"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.RouterGroup) {
	admin := router.Group("/admin")
	{
		admin.GET("/home", home.Index)
		admin.GET("/games", games.Index)
	}
}
