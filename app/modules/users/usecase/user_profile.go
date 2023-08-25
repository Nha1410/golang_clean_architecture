package usecase

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/modules/auth"
	"github.com/team2/real_api/app/models"
)

func (u UserUseCase) UserProfile(ctx *fiber.Ctx) (*models.UserResponse, error) {
	tokenString := strings.TrimPrefix(ctx.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		return nil, errors.New("unauthorized request")
	}

	// Verify the token
	userID, err := auth.VerifyToken(tokenString)
	if err != nil {
		return nil, errors.New("unauthorized request")
	}

	// Use the user's ID to fetch the user's profile
	userProfile, err := u.userRepo.GetUserProfile(userID)

	if err != nil {
		return nil, errors.New("unauthorized request")
	}

	return models.FilterUserRecord(userProfile), nil
}
