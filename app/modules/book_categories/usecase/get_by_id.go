package usecase

import (
	"errors"
	"github.com/team2/real_api/app/models"
)

func (u BookCategoryUseCase) GetBookCategoryByID(id int) (*models.BookCategoryResponse, error) {
	info, err := u.bookCategoryRepo.GetByID(id)

	if err != nil {
		return nil, errors.New("Book Category not found!")
	}

	return models.FilterBookCategoryRecord(info), nil
}