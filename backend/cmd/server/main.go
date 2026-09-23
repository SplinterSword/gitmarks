package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SplinterSword/gitmarks/backend/internal/bookmarks"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	bookmarks.RegisterRoutes(router)

	fmt.Println("Server listening on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
