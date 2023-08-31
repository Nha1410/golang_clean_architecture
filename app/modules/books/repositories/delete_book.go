package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r BookRepo) DeleteBook(bookID int) error {
	var book *models.Book

	result := r.DB.Table(models.Book{}.TableName()).Where("id = ?", bookID).First(&book).Delete(&book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

