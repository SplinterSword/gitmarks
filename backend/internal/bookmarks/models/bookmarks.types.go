package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Bookmark struct {
	ID    bson.ObjectID `bson:"_id,omitempty" json:"id"`
	URL   string        `bson:"url" json:"url"`
	Marks map[int]string `bson:"marks" json:"marks"`
}

type SaveBookmarkRequest struct {
	URL  string `json:"url"`
	File string `json:"file"`
	Mark int `json:"mark"`
}

type DeleteBookmarkRequest struct {
	URL  string `json:"url"`
	Mark int `json:"mark"`
}

type GetBookmarksRequest struct {
	URL  string `json:"url"`
}

type SaveBookmarkResponse struct {
	Message string `json:"message"`
	URL  string `json:"url"`
	File string `json:"file"`
	Mark int `json:"mark"`
}

type DeleteBookmarkResponse struct {
	Message string `json:"message"`
	URL  string `json:"url"`
	Mark int `json:"mark"`
}

type GetBookmarksResponse struct {
	Message string `json:"message"`
	Bookmarks map[string]map[int]string `json:"bookmarks"`
}
