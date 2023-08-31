package repository

import (
	"github.com/team2/real_api/app/models"
)

func(r BookCategoryRepo) DeleteBookCategory(bookCategory *models.BookCategory) error {
	result := r.DB.Table(models.BookCategory{}.TableName()).Delete(&bookCategory)
	if result.Error != nil {
		return result.Error
	}

	return nil
}