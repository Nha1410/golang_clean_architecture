package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	bookCategoryRepo "github.com/team2/real_api/app/modules/book_categories/repositories"
)

type UseCase interface {
	GetList() ([]*models.BookCategoryResponse, error)
	GetBookCategoryByID(id int) (*models.BookCategoryResponse, error)
	CreateBookCategory(ctx *fiber.Ctx, payload *models.BookCategoryInput) (*models.BookCategoryResponse, map[string]string)
	UpdateBookCategory(ctx *fiber.Ctx, id int, payload *models.BookCategoryInput) (*models.BookCategoryResponse, map[string]string)
	DeleteBookCategory(ctx *fiber.Ctx, id int) error
}


type BookCategoryUseCase struct {
	bookCategoryRepo bookCategoryRepo.BookCategoryRepoInterface
}

func NewBookCategoryUseCase(bookCategoryRepo bookCategoryRepo.BookCategoryRepoInterface) UseCase {
	return &BookCategoryUseCase{bookCategoryRepo: bookCategoryRepo}
}