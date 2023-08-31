package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (u BookCategoryUseCase) CreateBookCategory(ctx *fiber.Ctx, payload *models.BookCategoryInput) (*models.BookCategoryResponse, map[string]string) {
	createdBookCategory, err := u.bookCategoryRepo.CreateBookCategory(ctx, payload)

	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	return models.FilterBookCategoryRecord(createdBookCategory), nil
}