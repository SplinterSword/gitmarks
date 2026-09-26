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
		log.Fatal("Error loading .env")
	}

	utils.GlobalConfig["MONGODB_URI"] = os.Getenv("MONGODB_URI")
	utils.GlobalConfig["MONGODB_DATABASE"] = os.Getenv("MONGODB_DATABASE")

	router := httprouter.New()

	bookmarks.RegisterRoutes(router)

	if err := utils.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
