package main

import (
	"api/url-shorter/internal/api/routers"
	"api/url-shorter/internal/db"
	"fmt"
	"os"
)

func main() {
	uri := os.Getenv("MONGOURI")
	db.ConnectDB(uri)

	router := routers.SetupRouters()

	fmt.Printf("PORT: %v", os.Getenv("PORT"))
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
