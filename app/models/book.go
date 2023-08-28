package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Name        		string 		`gorm:"type:varchar(255)" json:"name"`
	Image       		string 		`json:"image"`
	Author      		string 		`gorm:"type:varchar(255)" json:"author"`
	PublicDate  		time.Time `json:"public_date"`
	Description 		string    `json:"description"`

	BookCategoryID  uint			`json:"book_category_id"`
	BookCategory 		BookCategory

	UserID      		uint			`json:"user_id"`
	User 						User

}

func (Book) TableName() string {
	return "books"
}
