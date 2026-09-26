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

	field := fmt.Sprintf("marks.%d", mark)

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

// DeleteFile deletes a file to a URL's bookmark document.
// If the URL doesn't exist, MongoDB raises an error.
func (r *BookmarkRepository) DeleteFile(
	ctx context.Context,
	url string,
	mark int,
) error {
	filter := bson.M{
		"url": url,
	}

	field := fmt.Sprintf("marks.%d", mark)

	update := bson.M{
		"$unset": bson.M{
			field: "",
		},
	}

	_, err := r.collection.UpdateOne(
		ctx,
		filter,
		update,
	)

	return err
}

// GetFiles get all the files to a URL's bookmark document.
// If the URL doesn't exist, MongoDB raises an error.
func (r *BookmarkRepository) GetFiles(
	ctx context.Context,
	url string,
) (map[int]string, error) {

	filter := bson.M{
		"url": url,
	}

	var bookmark Bookmark

	err := r.collection.FindOne(
		ctx,
		filter,
	).Decode(&bookmark)

	if err != nil {
		return nil, err
	}

	return bookmark.Marks, nil
}
