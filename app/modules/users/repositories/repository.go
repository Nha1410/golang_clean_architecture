package repository

import (
	"github.com/team2/real_api/app/models"
	"gorm.io/gorm"
)

type UserRepoInterface interface {
	ListUser() ([]*models.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}
