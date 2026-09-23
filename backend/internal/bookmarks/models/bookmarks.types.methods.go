package models

import (
	"context"
	"errors"

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
) error {

	filter := bson.M{
		"url": url,
	}

	update := bson.M{
		"$addToSet": bson.M{
			"files": file,
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

// GetBookmark returns the bookmark document for a URL.
func (r *BookmarkRepository) GetBookmark(
	ctx context.Context,
	url string,
) (*Bookmark, error) {

	var bookmark Bookmark

	err := r.collection.FindOne(
		ctx,
		bson.M{"url": url},
	).Decode(&bookmark)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, err
	}

	return &bookmark, nil
}

// GetFiles returns all files associated with a URL.
func (r *BookmarkRepository) GetFiles(
	ctx context.Context,
	url string,
) ([]string, error) {

	bookmark, err := r.GetBookmark(ctx, url)

	if err != nil {
		return nil, err
	}

	if bookmark == nil {
		return []string{}, nil
	}

	return bookmark.Files, nil
}

// RemoveFile removes a specific file from a URL.
func (r *BookmarkRepository) RemoveFile(
	ctx context.Context,
	url string,
	file string,
) error {

	filter := bson.M{
		"url": url,
	}

	update := bson.M{
		"$pull": bson.M{
			"files": file,
		},
	}

	_, err := r.collection.UpdateOne(
		ctx,
		filter,
		update,
	)

	return err
}

// HasFile checks whether a specific file exists for a URL.
func (r *BookmarkRepository) HasFile(
	ctx context.Context,
	url string,
	file string,
) (bool, error) {

	filter := bson.M{
		"url":   url,
		"files": file,
	}

	err := r.collection.FindOne(
		ctx,
		filter,
	).Err()

	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteBookmark removes the entire URL document.
func (r *BookmarkRepository) DeleteBookmark(
	ctx context.Context,
	url string,
) error {

	_, err := r.collection.DeleteOne(
		ctx,
		bson.M{"url": url},
	)

	return err
}
