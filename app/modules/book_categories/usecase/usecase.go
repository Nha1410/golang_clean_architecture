package usecase

import (
	"github.com/team2/real_api/app/models"
	bookCategoryRepo "github.com/team2/real_api/app/modules/book_categories/repositories"
)

type UseCase interface {
	GetList() ([]models.BookCategory, error)
	GetBookCategoryByID(id uint) (*models.BookCategory, error)
	CreateBookCategory(bookCategory *models.BookCategory) error
	UpdateBookCategory(bookCategory *models.BookCategory) error
	DeleteBookCategory(id uint) error
 }


type BookCategoryUseCase struct {
	bookCategoryRepo bookCategoryRepo.BookCategoryRepoInterface
}

func NewBookCategoryUseCase(bookCategoryRepo bookCategoryRepo.BookCategoryRepoInterface) UseCase {
	return &BookCategoryUseCase{bookCategoryRepo: bookCategoryRepo}
}