package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SplinterSword/gitmarks/backend/internal/bookmarks"
	"github.com/SplinterSword/gitmarks/backend/internal/utils"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	utils.GlobalConfig["MONGODB_URI"] = os.Getenv("MONGODB_URI")
	utils.GlobalConfig["MONGODB_DATABASE"] = os.Getenv("MONGODB_DATABASE")

	if utils.GlobalConfig["MONGODB_URI"] == "" || utils.GlobalConfig["MONGODB_DATABASE"] == "" {
		log.Fatal("MONGODB_URI and MONGODB_DATABASE must be set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()

	bookmarks.RegisterRoutes(router)

	if err := utils.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
