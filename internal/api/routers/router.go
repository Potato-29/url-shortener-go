package routers

import (
	"api/url-shorter/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.GET("/test", handlers.GetTestHandler)
	router.POST("/shorten", handlers.ShortenUrl)
	router.GET("/:id", handlers.GetShortenedURL)

	return router
}
