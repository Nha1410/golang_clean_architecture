package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	books "github.com/team2/real_api/app/modules/books/repositories"
)

type UseCase interface {
	AddBook(ctx *fiber.Ctx, payload *models.BookInput) (*models.BookResponse, map[string]string)
	EditBook(ctx *fiber.Ctx, bookID int, payload *models.BookInput) (*models.BookResponse, map[string]string)
	GetBook(bookID int) (*models.BookResponse, error)
	GetBooks() ([]*models.BookResponse, error)
	DeleteBook(ctx *fiber.Ctx, bookID int) error
}

type BookUseCase struct {
	bookRepo books.BookRepoInterface
}

func NewBookUseCase(bookRepo books.BookRepoInterface) UseCase {
	return &BookUseCase{bookRepo: bookRepo}
}
