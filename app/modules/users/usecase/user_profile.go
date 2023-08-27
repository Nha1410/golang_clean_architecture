package usecase

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (u UserUseCase) UserProfile(ctx *fiber.Ctx) (*models.UserResponse, error) {
	userID := ctx.Locals("userID").(int)
	userProfile, err := u.userRepo.GetUserProfile(userID)

	if err != nil {
		return nil, errors.New("unauthorized request")
	}

	return models.FilterUserRecord(userProfile), nil
}
