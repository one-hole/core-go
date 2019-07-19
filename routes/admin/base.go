package admin

import (
	"core-go/controllers/admin/battles"
	"core-go/controllers/admin/games"
	"core-go/controllers/home"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.RouterGroup) {
	admin := router.Group("/admin")
	{
		admin.GET("/games", games.Index)
		admin.GET("/battles", battles.Index)
	}
}
