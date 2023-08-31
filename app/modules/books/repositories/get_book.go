package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r BookRepo) GetBookByID(bookID int) (*models.Book, error) {
	var book *models.Book

	result := r.DB.Table(models.Book{}.TableName()).Preload("User").Preload("Category").Where("id = ?", bookID).First(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}
