package repository

import (
	"github.com/team2/real_api/app/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	GetBooks(userID int) ([]models.Book, error)
	GetUserProfile(userID int) (*models.User, error)
	CreateUser(data *models.SignUpInput) (*models.User, error)
	CheckEmailExists(email string) (bool)
	FindUserByEmail(email string) (*models.User, error)
	DeleteUser(user *models.User) error
	GetBookCategories(userID int) ([]models.BookCategory, error) 
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}
