package bookmarks

import "github.com/julienschmidt/httprouter"

func RegisterRoutes(router *httprouter.Router) {
	router.POST("/bookmark/save", SaveBookmark)
}
