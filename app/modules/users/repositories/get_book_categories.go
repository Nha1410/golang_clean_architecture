package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r UserRepo) GetBookCategories(userID int) ([]models.BookCategory, error) {
	var bookCategories []models.BookCategory

	result := r.DB.Table(models.BookCategory{}.TableName()).Where("user_id = ?", userID).Find(&bookCategories)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookCategories, nil
}