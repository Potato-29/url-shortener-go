package services

import (
	"api/url-shorter/internal/db"
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlHash struct {
	ID        primitive.ObjectID `json:"_id"`
	BaseUrl   string             `json:"base-url"`
	Hash      string             `json:"url-hash"`
	CreatedAt time.Time          `json:"createdAt"`
}

var temp_data = []string{"1", "2", "3"}

func TestData() ([]string, error) {
	return temp_data, nil
}

func InsertUrlDocument(BaseUrl string, urlHash string) (string, error) {

	doc := UrlHash{
		ID:        primitive.NewObjectID(),
		BaseUrl:   BaseUrl,
		Hash:      urlHash,
		CreatedAt: time.Now(),
	}

	_, insertErr := db.GetCollection("url-hashes").InsertOne(context.TODO(), doc)
	if insertErr != nil {
		return "", errors.New("Failed to save shortened URL")
	}
	return os.Getenv("BASEURL") + urlHash, nil
}

func GetUrlDocumentByID(id string) (string, error) {
	var result UrlHash

	err := db.GetCollection("url-hashes").FindOne(context.TODO(), bson.M{"hash": id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("document not found")
		}
		return "", fmt.Errorf("failed to fetch document: %v", err)
	}

	redirectUrl := result.BaseUrl

	return redirectUrl, nil
}
