package handlers

import (
	"api/url-shorter/internal/api/services"
	"api/url-shorter/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTestHandler(c *gin.Context) {
	res, err := services.TestData()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	c.IndentedJSON(http.StatusOK, res)
}

func ShortenUrl(c *gin.Context) {
	var input struct {
		BaseUrl string `json:"base-url"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	urlHash, hashErr := utils.GenerateRandomBase64Hash(8)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
	}

	result, err := services.InsertUrlDocument(input.BaseUrl, urlHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "URL shortened successfully",
		"data":    result,
	})
}

func GetShortenedURL(c *gin.Context) {
	id := c.Param("id")

	// Fetch the shortened URL document by ID
	urlDoc, err := services.GetUrlDocumentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": urlDoc})
}
