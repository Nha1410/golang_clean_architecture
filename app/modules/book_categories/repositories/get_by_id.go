package repository

import "github.com/team2/real_api/app/models"

func (r BookCategoryRepo) GetByID(id int) (*models.BookCategory, error) {
	var bookCategory *models.BookCategory

	result := r.DB.Table(models.BookCategory{}.TableName()).Preload("User").Preload("Books").Where("id = ? AND deleted_at IS NULL", id).First(&bookCategory)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookCategory, nil
}