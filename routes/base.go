package routes

import (
	"core-go/routes/api"
	"net/http"
	"os"
	"time"

	"core-go/controllers/books"
	"core-go/routes/admin"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := routers()
	s := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func routers() *gin.Engine {
	gin.SetMode(os.Getenv("GO_ENV"))
	router := gin.Default()
	root := router.Group("")
	{
		admin.ApplyRoutes(root)
		api.ApplyRoutes(root)

		root.GET("/books", books.Index)
		root.GET("/books/:id", books.Show)
	}
	return router
}
