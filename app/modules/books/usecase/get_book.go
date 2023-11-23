package usecase

import (
	"errors"
	"github.com/team2/real_api/app/models"
)

func (u BookUseCase) GetBook(bookID int) (*models.BookResponse, error) {
	bookInfo, err := u.bookRepo.GetBookByID(bookID)

	if err != nil {
		return nil, errors.New("Book not found!")
	}

	return models.FilterBookRecord(bookInfo), nil
}
