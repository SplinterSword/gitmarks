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
