package repository

import "github.com/team2/real_api/app/models"

func (r BookCategoryRepo) GetList() ([]*models.BookCategory, error) {
	var bookCategories []*models.BookCategory

	result := r.DB.Table(models.BookCategory{}.TableName()).Preload("User").Preload("Books").Find(&bookCategories)

	if result.Error != nil {
		return nil, result.Error
	}

	return bookCategories, nil
}
