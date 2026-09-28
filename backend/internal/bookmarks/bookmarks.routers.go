package bookmarks

import "github.com/julienschmidt/httprouter"

func RegisterRoutes(router *httprouter.Router) {
	router.POST("/bookmarks/save", SaveBookmarkController)
	router.DELETE("/bookmarks/delete", DeleteBookmarkController)
	router.GET("/bookmarks/get", GetBookmarksController)
}
