package models

type BookmarkCollections struct {
	Collections map[string][]string `json:"collections"`
}

type SaveBookmarkRequest struct {
	URL   string `json:"url"`
	File  string `json:"file"`
}

