package api

import (
	"core-go/controllers/api/battles"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(router *gin.RouterGroup) {
	root := router.Group("")
	{
		root.GET("/schedules", battles.Index())
	}
}
