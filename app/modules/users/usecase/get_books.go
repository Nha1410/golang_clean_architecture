package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (u UserUseCase) GetBooks(ctx *fiber.Ctx) ([]*models.OnlyBookResponse, error) {
	userID := ctx.Locals("userID").(int)
	books, err := u.userRepo.GetBooks(userID)

	if err != nil {
		return nil, err
	}

	return models.FilterListBookOnlyRecord(books), nil
}
