package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func(u UserUseCase) GetBookCategories(ctx *fiber.Ctx) ([]*models.OnlyBookCategoryResponse, error) {
	
	userID := ctx.Locals("userID").(int)
	bookCategories, err := u.userRepo.GetBookCategories(userID)

	if err != nil {
		return nil, err
	}

	return models.FilterListBookCategoryOnlyRecord(bookCategories), nil
} 