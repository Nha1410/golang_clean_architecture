package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	users "github.com/team2/real_api/app/modules/users/repositories"
)

type UseCase interface {
	GetBooks(ctx *fiber.Ctx) ([]*models.OnlyBookResponse, error)
	UserProfile(ctx *fiber.Ctx) (*models.UserResponse, error)
	SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.SignUpResponse, map[string]string)
	SignInUser(email, password string) (string, error)	
	DeleteUser(ctx *fiber.Ctx, userID int) error
	GetBookCategories(ctx *fiber.Ctx) ([]*models.OnlyBookCategoryResponse, error)
}

type UserUseCase struct {
	userRepo users.UserRepoInterface
}

func NewUserUseCase(userRepo users.UserRepoInterface) UseCase {
	return &UserUseCase{userRepo: userRepo}
}
