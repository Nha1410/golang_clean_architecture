package handlers

import (
	book "github.com/team2/real_api/app/modules/books/usecase"
)

type BookHandlers struct {
	bookUseCase book.UseCase
}

func NewBookHandlers(bookUseCase book.UseCase) *BookHandlers {
	return &BookHandlers{bookUseCase: bookUseCase}
}
