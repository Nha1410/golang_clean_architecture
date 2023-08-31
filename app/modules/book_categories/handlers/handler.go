package handlers

import (
	bookCategory "github.com/team2/real_api/app/modules/book_categories/usecase"
)

type BookCategoryHandlers struct {
	BookCategoryUseCase bookCategory.UseCase
}

func NewBookCategoryHandlers(BookCategory bookCategory.UseCase) *BookCategoryHandlers { 
	return &BookCategoryHandlers{BookCategoryUseCase: bookCategory.BookCategoryUseCase{}}
}