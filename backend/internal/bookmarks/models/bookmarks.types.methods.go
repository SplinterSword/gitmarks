package models

import (
	"errors"
	"slices"
)


// Constructor
func NewBookmarkCollections() *BookmarkCollections {
	return &BookmarkCollections{
		Collections: make(map[string][]string),
	}
}


// Add Bookmark
func (b *BookmarkCollections) AddBookmark(collection string, bookmarkID string) {
	b.Collections[collection] = append(b.Collections[collection], bookmarkID)
}

// Get Bookmarks
func (b *BookmarkCollections) GetBookmarks(collection string) []string {
	return b.Collections[collection]
}

// Delete bookmark
func (b *BookmarkCollections) RemoveBookmark(collection string, bookmarkID string) error {
	bookmarks, found := b.Collections[collection]
	if found != true {
		return errors.New("Bookmark Collection not found")
	}

	for i, id := range bookmarks{
		if id == bookmarkID {
			b.Collections[collection] = append(bookmarks[:i], bookmarks[i+1:]...)
			return nil
		}
	}

	return errors.New("bookmarkID: " + bookmarkID + " not found in the Bookmark Collection")
}

// Has Bookmark
func (b *BookmarkCollections) HasBookmark(collection string, bookmarkID string) bool {
	bookmarks := b.Collections[collection]

	found := slices.Contains(bookmarks, bookmarkID)

	return found
}
