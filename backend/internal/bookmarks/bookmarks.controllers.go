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
	mark := schema.Mark
	context := r.Context()

	if err := SaveBookmark(context, url, file, mark); err != nil {
		utils.SendError(err, http.StatusInternalServerError, w)
	}

	response := models.SaveBookmarkResponse{
		Message: "Successfully Added to the cloud",
		URL:     url,
		File:    file,
		Mark:    mark,
	}
	utils.SendJson(response, w, r)
}

func DeleteBookmarkController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var schema models.DeleteBookmarkRequest

	if err := utils.ReceiveJson(&schema, w, r); err != nil {
		return
	}

	url := schema.URL
	mark := schema.Mark
	context := r.Context()

	if err := DeleteBookmark(context, url, mark); err != nil {
		utils.SendError(err, http.StatusInternalServerError, w)
	}

	response := models.DeleteBookmarkResponse{
		Message: "Successfully Deleted from the cloud",
		URL:     url,
		Mark:    mark,
	}
	utils.SendJson(response, w, r)
}

func GetBookmarksController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var schema models.GetBookmarksRequest

	if err := utils.ReceiveJson(&schema, w, r); err != nil {
		return
	}

	url := schema.URL
	context := r.Context()

	marks, err := GetBookmarks(context, url)
	if err != nil {
		utils.SendError(err, http.StatusInternalServerError, w)
	}

	response := models.GetBookmarksResponse{
		Message: "Bookmarks retreived successfully",
		Bookmarks: map[string]map[int]string {
			url: marks,
		},
	}

	utils.SendJson(response, w, r)
	
}
