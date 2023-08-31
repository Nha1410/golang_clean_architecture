package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (u BookUseCase) EditBook(ctx *fiber.Ctx, bookID int, payload *models.BookInput) (*models.BookResponse, map[string]string) {
	book, err := u.bookRepo.GetBookByID(bookID)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	updatedBook, updateErr := u.bookRepo.UpdateBook(book, payload)
	if updateErr != nil {
		return nil, map[string]string{"error": updateErr.Error()}
	}

	return models.FilterBookRecord(updatedBook), nil
}
