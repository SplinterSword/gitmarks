package utils

import (
	"github.com/SplinterSword/gitmarks/backend/internal/bookmarks/models"
)

type Storage struct {
	BookmarkCollections *models.BookmarkCollections `json:"bookmark-collection"`
}

var GlobalStorage = &Storage {
	BookmarkCollections: models.NewBookmarkCollections(),
}
