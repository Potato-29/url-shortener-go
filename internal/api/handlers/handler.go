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

func GetTestHandler(c *gin.Context) {
	res, err := services.TestData()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func ShortenUrl(c *gin.Context) {
	var input struct {
		BaseUrl string `json:"base-url"`
	}
	var duplicate services.UrlHash
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	urlHash, hashErr := utils.GenerateRandomBase64Hash(8)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
		return
	}

	err := db.GetCollection("url-hashes").FindOne(context.TODO(), bson.M{"hash": urlHash}).Decode(&duplicate)

	if err == nil {
		ShortenUrl(c)
	}

	result, err := services.InsertUrlDocument(input.BaseUrl, urlHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL shortened successfully",
		"data":    result,
	})
	return
}

func ShortenUrlWithAlias(c *gin.Context) {
	var input struct {
		BaseUrl string `json:"base-url"`
		Alias   string `json:"alias"`
	}
	var duplicate services.UrlHash
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := db.GetCollection("url-hashes").FindOne(context.TODO(), bson.M{"hash": input.Alias}).Decode(&duplicate)
	fmt.Printf("MY ERR: %v \n", err)
	if err == nil {
		c.JSON(http.StatusConflict, "Please use a unique alias!")
		return
	}
	fmt.Printf("aliassss: %v \n", input.Alias)
	result, err := services.InsertUrlDocument(input.BaseUrl, input.Alias)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL shortened successfully",
		"data":    result,
	})
}

func GetShortenedURL(c *gin.Context) {
	id := c.Param("id")

	// Fetch the shortened URL document by ID
	redirectUrl, err := services.GetUrlDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.Redirect(301, redirectUrl)
}
