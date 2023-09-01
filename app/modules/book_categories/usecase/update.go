package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (u BookCategoryUseCase) UpdateBookCategory(ctx *fiber.Ctx, id int, payload *models.BookCategoryInput) (*models.BookCategoryResponse, map[string]string) {
	bookCategory, err := u.bookCategoryRepo.GetByID(id)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	updatedBookCategory, updateErr := u.bookCategoryRepo.UpdateBookCategory(ctx, bookCategory, payload)
	if updateErr != nil {
		return nil, map[string]string{"error": updateErr.Error()}
	}

	return models.FilterBookCategoryRecord(updatedBookCategory), nil
}