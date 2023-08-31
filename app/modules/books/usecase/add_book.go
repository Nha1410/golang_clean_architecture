package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (u BookUseCase) AddBook(ctx *fiber.Ctx, payload *models.BookInput) (*models.BookResponse, map[string]string) {
	createdBook, err := u.bookRepo.CreateBook(payload)

	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	return models.FilterBookRecord(createdBook), nil
}
