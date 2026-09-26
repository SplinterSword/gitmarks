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
