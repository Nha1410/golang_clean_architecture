package repository

import (
	"github.com/team2/real_api/app/models"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

type BookRepoInterface interface {
	CreateBook(ctx *fiber.Ctx, data *models.BookInput) (*models.Book, error)
	UpdateBook(ctx *fiber.Ctx, book *models.Book, payload *models.BookInput) (*models.Book, error)
	GetBookByID(bookID int) (*models.Book, error)
	GetBooks() ([]*models.Book, error)
	DeleteBook(bookID int) error
}

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{DB: db}
}
