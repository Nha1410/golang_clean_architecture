package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r BookRepo) GetBooks() ([]*models.Book, error) {
	var books []*models.Book

	result := r.DB.Table(models.Book{}.TableName()).Preload("User").Preload("BookCategory").Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

