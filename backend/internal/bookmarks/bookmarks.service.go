package bookmarks

import (
	"context"

	"github.com/SplinterSword/gitmarks/backend/internal/utils"
)

func SaveBookmark(ctx context.Context, url string, file string, mark int) error {
	if err := utils.GlobalStorage.BookmarkRepository.AddFile(ctx, url, file, mark); err != nil {
		return err
	}

	return nil
}

func DeleteBookmark(ctx context.Context, url string, mark int) error {
	if err := utils.GlobalStorage.BookmarkRepository.DeleteFile(ctx, url, mark); err != nil {
		return err
	}

	return nil
}

func GetBookmarks(ctx context.Context, url string) (map[int]string, error) {
	marks, err := utils.GlobalStorage.BookmarkRepository.GetFiles(ctx, url)
	if err != nil {
		return make(map[int]string),err
	}

	return marks,nil
}
