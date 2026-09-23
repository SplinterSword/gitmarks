package utils

import (
	"github.com/SplinterSword/gitmarks/backend/internal/bookmarks/models"
)

type Storage struct {
	BookmarkCollections *models.BookmarkRepository
	MongoDB *MongoDB
}

var GlobalStorage = &Storage{}

func ConnectDatabase() error {
	mongoDB, err := Connect(GlobalConfig["MONGODB_URI"], GlobalConfig["MONGODB_DATABASE"])
	
	if err != nil {
		return err
	}

	GlobalStorage.MongoDB = mongoDB

	GlobalStorage.BookmarkCollections = models.NewBookmarkRepository(mongoDB.Database)
	return nil
}
