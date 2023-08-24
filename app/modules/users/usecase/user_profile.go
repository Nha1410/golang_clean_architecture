package usecase

import (
	"errors"
	"github.com/team2/real_api/app/models"
	"github.com/team2/real_api/app/auth"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func (u *UserUseCase) UserProfile(ctx *fiber.Ctx) (*models.UserResponse, error) {
	tokenString := strings.TrimPrefix(ctx.Get("Authorization"), "Bearer ")
	if tokenString == "" {
		return nil, errors.New("Unauthorized request")
	}

	// Verify the token
	userID, err := auth.VerifyToken(tokenString)
	if err != nil {
		return nil, errors.New("Unauthorized request")
	}

	// Use the user's ID to fetch the user's profile
	userProfile, err := u.userRepo.GetUserProfile(userID)

	if err != nil {
		return nil, errors.New("Unauthorized request")
	}

	return models.FilterUserRecord(userProfile), nil
}
