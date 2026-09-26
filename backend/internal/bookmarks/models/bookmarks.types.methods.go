package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BookmarkRepository struct {
	collection *mongo.Collection
}

func NewBookmarkRepository(db *mongo.Database) *BookmarkRepository {
	return &BookmarkRepository{
		collection: db.Collection("bookmarks"),
	}
}

// AddFile adds a file to a URL's bookmark document.
// If the URL doesn't exist, MongoDB creates the document.
func (r *BookmarkRepository) AddFile(
	ctx context.Context,
	url string,
	file string,
	mark int,
) error {
	filter := bson.M{
		"url": url,
	}

	field := fmt.Sprintf("mark.%d", mark)

	update := bson.M{
		"$set": bson.M{
			field: file,
		},
	}

	_, err := r.collection.UpdateOne(
		ctx,
		filter,
		update,
		options.UpdateOne().SetUpsert(true),
	)

	return err
}

