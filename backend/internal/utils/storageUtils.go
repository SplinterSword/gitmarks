package utils

import (
	"github.com/SplinterSword/gitmarks/backend/internal/bookmarks/models"
)

type Storage struct {
	BookmarkRepository *models.BookmarkRepository
	MongoDB *MongoDB
}

var GlobalStorage = &Storage{}

func ConnectDatabase() error {
	mongoDB, err := Connect(GlobalConfig["MONGODB_URI"], GlobalConfig["MONGODB_DATABASE"])
	
	if err != nil {
		return err
	}

	GlobalStorage.MongoDB = mongoDB

	GlobalStorage.BookmarkRepository= models.NewBookmarkRepository(mongoDB.Database)
	return nil
}
