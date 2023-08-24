package usecase

import (
	"github.com/team2/real_api/app/models"
	users "github.com/team2/real_api/app/modules/users/repositories"
	"github.com/gofiber/fiber/v2"
)

type UseCase interface {
	ListUser(ctx *fiber.Ctx) ([]*models.UserResponse, error)
}

type UserUseCase struct {
	userRepo users.UserRepoInterface
}

func NewUserUseCase(userRepo users.UserRepoInterface) UseCase {
	return &UserUseCase{userRepo: userRepo}
}
