package bookmarks

import "github.com/julienschmidt/httprouter"

func RegisterRoutes(router *httprouter.Router) {
	router.POST("/bookmark/save", SaveBookmarkController)
	router.DELETE("/bookmark/delete", DeleteBookmarkController)
	router.GET("/bookmark/get", GetBookmarksController)
}
