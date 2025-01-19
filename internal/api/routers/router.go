package routers

import (
	"api/url-shorter/internal/api/handlers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Working Directory:", wd)

	templatePath := wd + "/internal/templates/*.html"

	router.LoadHTMLGlob(templatePath)

	router.GET("/", handlers.RenderHome)
	router.GET("/:id", handlers.GetShortenedURL)

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/test", handlers.GetTestHandler)
		apiGroup.POST("/shorten", handlers.ShortenUrl)
	}

	return router
}
