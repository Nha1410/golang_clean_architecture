package usecase

import (
	"github.com/team2/real_api/app/models"
	"github.com/gofiber/fiber/v2"
)

func (u BookCategoryUseCase) UpdateBookCategory(ctx *fiber.Ctx, id int, payload *models.BookCategoryInput) (*models.BookCategoryResponse, map[string]string) {
	book, err := u.bookCategoryRepo.GetByID(id)
	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	updatedBookCategory, updateErr := u.bookCategoryRepo.UpdateBookCategory(ctx, book, payload)
	if updateErr != nil {
		return nil, map[string]string{"error": updateErr.Error()}
	}

	return models.FilterBookCategoryRecord(updatedBookCategory), nil
}