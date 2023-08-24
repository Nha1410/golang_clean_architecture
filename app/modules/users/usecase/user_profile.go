package usecase

import (
	"errors"
	"github.com/team2/real_api/app/models"
	"github.com/team2/real_api/app/auth"
	"github.com/gofiber/fiber/v2"
)

func (u *UserUseCase) UserProfile(ctx *fiber.Ctx) (*models.UserResponse, error) {
	tokenString := ctx.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("missing authorization token")
	}

	// Verify the token
	userID, err := auth.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	// Use the user's ID to fetch the user's profile
	userProfile, err := u.userRepo.GetUserProfile(userID)

	if err != nil {
		return nil, err
	}

	return models.FilterUserRecord(userProfile), nil
}
