package bookmarks

import (
	"net/http"

	"github.com/SplinterSword/gitmarks/backend/internal/bookmarks/models"
	"github.com/SplinterSword/gitmarks/backend/internal/utils"
	"github.com/julienschmidt/httprouter"
)

func SaveBookmarkController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var schema models.SaveBookmarkRequest

	if err := utils.ReceiveJson(&schema, w, r); err != nil {
		return
	}

	url := schema.URL
	file := schema.File

	utils.GlobalStorage.BookmarkCollections.AddBookmark(url, file)

	utils.SendJson(utils.GlobalStorage.BookmarkCollections, w, r)
}
