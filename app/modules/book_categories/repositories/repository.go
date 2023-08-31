package repository

import (
	"github.com/team2/real_api/app/models"
	"gorm.io/gorm"
)

type BookCategoryRepoInterface interface {
	GetList() ([]models.BookCategory, error)
	GetByID(id uint) (*models.BookCategory, error)
	Create(bookCategory *models.BookCategory) (*models.BookCategory, error)
	Update(bookCategory *models.BookCategory) (*models.BookCategory, error)
	Delete(id uint) error
}

type BookCategoryRepo struct {
	DB *gorm.DB
}

func NewBookCategoryRepo(db *gorm.DB) *BookCategoryRepo {
	return &BookCategoryRepo{DB: db}
}