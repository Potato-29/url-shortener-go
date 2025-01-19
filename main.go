package main

import (
	"api/url-shorter/internal/api/routers"
	"api/url-shorter/internal/db"
	"fmt"
	"os"
)

func main() {
	uri := os.Getenv("MONGO_URI")
	db.ConnectDB(uri)

	router := routers.SetupRouters()

	router.Run(fmt.Sprintf("%s", os.Getenv("PORT")))
}
