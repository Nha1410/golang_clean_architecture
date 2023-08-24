package usecase

import (
	"github.com/team2/real_api/app/models"
	"github.com/gofiber/fiber/v2"
)

func (u *UserUseCase) ListUser(ctx *fiber.Ctx) ([]*models.UserResponse, error) {
	users, err := u.userRepo.ListUser()

	if err != nil {
		return nil, err
	}

	var userResponses []*models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, models.FilterUserRecord(user))
	}

	return userResponses, nil
}