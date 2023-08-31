package models

import (
	"gorm.io/gorm"
)

type BookCategory struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	User        User
	Books       []Book
}

func (BookCategory) TableName() string {
	return "book_categories"
}

type BookCategoryInfoResponse struct {
	ID          uint 		`json:"id"`
	Name        string 	`gorm:"type:varchar(255)" json:"name"`
	Image       string 	`json:"image"`
	Description string  `json:"description"`
}

func FilterBookCategoryInfoRecord(bookCategory *BookCategory) *BookCategoryInfoResponse {
	return &BookCategoryInfoResponse{
		ID: bookCategory.ID,
		Name: bookCategory.Name,
		Image: bookCategory.Image,
		Description: bookCategory.Description,
	}
}
