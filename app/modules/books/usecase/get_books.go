package usecase

import (
	"github.com/team2/real_api/app/models"
)

func (u BookUseCase) GetBooks() ([]*models.BookResponse, error) {
	books, err := u.bookRepo.GetBooks()
	if err != nil {
		return nil, err
	}

	var bookResponses []*models.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, models.FilterBookRecord(book))
	}


	return bookResponses, nil
}
