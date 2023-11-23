package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r UserRepo) GetBooks(userID int) ([]models.Book, error) {
	var books []models.Book

	result := r.DB.Table(models.Book{}.TableName()).Where("user_id = ?", userID).Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}
