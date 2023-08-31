package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Name        		string 				`gorm:"type:varchar(255)" json:"name"`
	Image       		string 				`json:"image"`
	Author      		string 				`gorm:"type:varchar(255)" json:"author"`
	PublicDate  		time.Time 		`json:"public_date"`
	Description 		string    		`json:"description"`

	BookCategoryID  uint					`json:"book_category_id"`
	BookCategory 		BookCategory

	UserID      		uint					`json:"user_id"`
	User 						User

}

func (Book) TableName() string {
	return "books"
}

type BookInput struct {
	Name           string `json:"name" validate:"required"`
	Image          string `json:"image" validate:"required"`
	Author         string `json:"author" validate:"required"`
	PublicDate     string `json:"public_date" validate:"required"`
	Description    string `json:"description" validate:"required"`
	BookCategoryID uint 	`json:"book_category_id" validate:"required"`
	UserID 				 uint 	`json:"user_id" validate:"required"`
}

type BookResponse struct {
	ID               uint 										`json:"id"`
	Name        		 string 									`gorm:"type:varchar(255)" json:"name"`
	Image       		 string 									`json:"image"`
	Author      		 string 									`gorm:"type:varchar(255)" json:"author"`
	PublicDate  		 time.Time 								`json:"public_date"`
	Description 		 string    								`json:"description"`
	BookCategoryID   uint											 `json:"book_category_id"`
	BookCategory     *BookCategoryInfoResponse `json:"book_category"`
	UserID      		 uint					 						 `json:"user_id"`
	User  					 *UserResponse 						 `json:"user"`
}

func FilterBookRecord(book *Book) *BookResponse {
	return &BookResponse{
		ID:        			book.ID,
		Name: 					book.Name,
		Image: 					book.Image,
		Author: 				book.Author,
		PublicDate: 		book.PublicDate,
		Description: 		book.Description,
		BookCategoryID: book.BookCategoryID,
		BookCategory:		FilterBookCategoryInfoRecord(&book.BookCategory),
		UserID: 				book.UserID,
		User: 					FilterUserRecord(&book.User),
	}
}

