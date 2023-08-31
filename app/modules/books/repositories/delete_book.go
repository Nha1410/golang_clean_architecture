package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r BookRepo) DeleteBook(book *models.Book) error {
	result := r.DB.Table(models.Book{}.TableName()).Delete(&book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

