package repository

import (
	"github.com/team2/real_api/app/models"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)

type BookCategoryRepoInterface interface {
	GetList() ([]*models.BookCategory, error)
	GetByID(id int) (*models.BookCategory, error)
	CreateBookCategory(ctx *fiber.Ctx, payload *models.BookCategoryInput) (*models.BookCategory, error)
	UpdateBookCategory(ctx *fiber.Ctx,bookCategory *models.BookCategory, payload *models.BookCategoryInput) (*models.BookCategory, error)
	DeleteBookCategory(bookCategory *models.BookCategory) error
}

type BookCategoryRepo struct {
	DB *gorm.DB
}

func NewBookCategoryRepo(db *gorm.DB) *BookCategoryRepo {
	return &BookCategoryRepo{DB: db}
}