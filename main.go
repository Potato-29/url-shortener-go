package main

import (
	"api/url-shorter/internal/api/routers"
	"api/url-shorter/internal/db"
)

func main() {
	uri := "mongodb+srv://testuser29:Test%40123@testcluster0.ppxbwit.mongodb.net/?retryWrites=true&w=majority&appName=testCluster0"
	db.ConnectDB(uri)

	router := routers.SetupRouters()

	router.Run("localhost:5000")
}
