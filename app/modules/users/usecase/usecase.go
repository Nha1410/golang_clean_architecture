package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	users "github.com/team2/real_api/app/modules/users/repositories"
)

type UseCase interface {
	UserProfile(ctx *fiber.Ctx) (*models.UserResponse, error)
	SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.SignUpResponse, map[string]string)
	SignInUser(email, password string) (string, error)	
}

type UserUseCase struct {
	userRepo users.UserRepoInterface
}

func NewUserUseCase(userRepo users.UserRepoInterface) UseCase {
	return &UserUseCase{userRepo: userRepo}
}
