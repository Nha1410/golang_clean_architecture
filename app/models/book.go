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

	BookCategoryID  int						`json:"book_category_id"`
	BookCategory 		BookCategory

	UserID      		int						`json:"user_id"`
	User 						User

}

func (Book) TableName() string {
	return "books"
}

type BookInput struct {
	Name           string `json:"name" validate:"required"`
	Image          string `json:"image"`
	Author         string `json:"author" validate:"required"`
	PublicDate     string `json:"public_date" validate:"required"`
	Description    string `json:"description" validate:"required"`
	BookCategoryID int 	`json:"book_category_id" validate:"required"`
}

type BookResponse struct {
	ID               uint 										`json:"id"`
	Name        		 string 									`gorm:"type:varchar(255)" json:"name"`
	Image       		 string 									`json:"image"`
	Author      		 string 									`gorm:"type:varchar(255)" json:"author"`
	PublicDate  		 time.Time 								`json:"public_date"`
	Description 		 string    								`json:"description"`
	BookCategoryID   int											 `json:"book_category_id"`
	BookCategory     *BookCategoryInfoResponse `json:"book_category"`
	UserID      		 int					 						 `json:"user_id"`
	User  					 *UserResponse 						 `json:"user"`
}

type onlyBookResponse struct {
	ID               uint 										`json:"id"`
	Name        		 string 									`gorm:"type:varchar(255)" json:"name"`
	Image       		 string 									`json:"image"`
	Author      		 string 									`gorm:"type:varchar(255)" json:"author"`
	PublicDate  		 time.Time 								`json:"public_date"`
	Description 		 string    								`json:"description"`
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

func FilterListBookRecord(books []Book) []*BookResponse {
	var bookResponses []*BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, FilterBookRecord(&book))
	}
	return bookResponses
}

func FilterOnlyBookRecord(book *Book) *onlyBookResponse {
	return &onlyBookResponse{
		ID:        			book.ID,
		Name: 					book.Name,
		Image: 					book.Image,
		Author: 				book.Author,
		PublicDate: 		book.PublicDate,
		Description: 		book.Description,
	}
}

func FilterListBookOnlyRecord(books []Book) []*onlyBookResponse {
	var bookResponses []*onlyBookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, FilterOnlyBookRecord(&book))
	}
	return bookResponses
}
