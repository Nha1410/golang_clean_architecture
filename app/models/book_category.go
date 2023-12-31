package models

import (
	"gorm.io/gorm"
)

type BookCategory struct {
	gorm.Model
	Name        string	`gorm:"type:varchar(255)" json:"name"`
	Image       string	`json:"image"`
	Description string	`json:"description"`
	UserID      int			`json:"user_id"`
	User        User
	Books       []Book	`gorm:"foreignKey:BookCategoryID"`
}

func (BookCategory) TableName() string {
	return "book_categories"
}

type BookCategoryInput struct {
	Name           string `json:"name" validate:"required"`
	Image          string `json:"image"`
	Description    string `json:"description" validate:"required"`
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

type BookCategoryResponse struct {
	ID          	uint 						 `json:"id"`
	Name        	string 					 `gorm:"type:varchar(255)" json:"name"`
	Image       	string 					 `json:"image"`
	Description 	string  		     `json:"description"`
	Books     		[]*OnlyBookResponse  `json:"books"`
	UserID      	int					 		 `json:"user_id"`
	User  				*UserResponse 	 `json:"user"`
}
type OnlyBookCategoryResponse struct {
	ID          	uint 				`json:"id"`
	Name        	string 				`gorm:"type:varchar(255)" json:"name"`
	Image       	string 				`json:"image"`
	Description 	string  		    `json:"description"`
}

func FilterBookCategoryRecord(bookCategory *BookCategory) *BookCategoryResponse {
	return &BookCategoryResponse{
		ID: 				 bookCategory.ID,
		Name: 			 bookCategory.Name,
		Image: 			 bookCategory.Image,
		Description: bookCategory.Description,
		Books:			 FilterListBookOnlyRecord(bookCategory.Books),
		UserID: 		 bookCategory.UserID,
		User: 			 FilterUserRecord(&bookCategory.User),
	}
}

func FilterOnlyBookCategoryRecord(bookCategory *BookCategory) *OnlyBookCategoryResponse {
	return &OnlyBookCategoryResponse{
		ID:        			bookCategory.ID,
		Name: 					bookCategory.Name,
		Image: 					bookCategory.Image,
		Description: 		bookCategory.Description,
	}
}


func FilterListBookCategoryOnlyRecord(bookCategoies []BookCategory) []*OnlyBookCategoryResponse  {
	var bookCategoryResponses []*OnlyBookCategoryResponse
	for _, bookCategory := range bookCategoies {
		bookCategoryResponses = append(bookCategoryResponses, FilterOnlyBookCategoryRecord(&bookCategory))
	}
	return bookCategoryResponses
}
