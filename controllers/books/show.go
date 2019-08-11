package books

import (
	serializer "core-go/serializers"

	"github.com/gin-gonic/gin"
	"github.com/gonrails/gonrails/serializers"
)

func Show(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": serializers.SingleSerializer(&serializer.BookSerializer{}, nil),
	})
}
