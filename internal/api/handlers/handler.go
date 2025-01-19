package handlers

import (
	"api/url-shorter/internal/api/services"
	"api/url-shorter/internal/db"
	"api/url-shorter/internal/pkg/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func RenderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "URL Shortener",
		"Message": "Test Page",
	})
}

func GetTestHandler(c *gin.Context) {
	res, err := services.TestData()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func ShortenUrl(c *gin.Context) {
	var duplicate services.UrlHash
	input := c.PostForm("base-url")
	alias := c.PostForm("alias")

	var urlHash string
	var hashErr error

	if alias == "" {
		fmt.Printf("null alias: %v", alias)
		urlHash, hashErr = utils.GenerateRandomBase64Hash(8)
		if hashErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
			return
		}
	} else {
		urlHash = alias
	}

	err := db.GetCollection("url-hashes").FindOne(context.TODO(), bson.M{"hash": urlHash}).Decode(&duplicate)

	if err == nil {
		ShortenUrl(c)
	}

	result, err := services.InsertUrlDocument(input, urlHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "response.html", gin.H{
		"ShortenedURL": result,
	})
	return
}

func GetShortenedURL(c *gin.Context) {
	id := c.Param("id")

	redirectUrl, err := services.GetUrlDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.Redirect(301, redirectUrl)
}
