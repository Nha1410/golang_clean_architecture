package repository

import (
	"github.com/team2/real_api/app/models"
	"time"
	"errors"
)

func (r BookRepo) UpdateBook(book *models.Book, payload *models.BookInput) (*models.Book, error) {
	var user models.User
	var category models.BookCategory
	err := r.DB.First(&user, payload.UserID).Error
	if err != nil {
			return nil, errors.New("user_id is not valid")
	}
	err = r.DB.First(&category, payload.BookCategoryID).Error
	if err != nil {
			return nil, errors.New("book_category_id is not valid")
	}

	book.Name = payload.Name
	book.Image = payload.Image
	book.Author = payload.Author
	parsedTime, err := time.Parse("01/02/2006", payload.PublicDate)
	if err != nil {
		return nil, errors.New("Please provide a valid date with format: mm/dd/yyyy")
	}
	book.PublicDate = parsedTime
	book.Description = payload.Description
	book.BookCategoryID = payload.BookCategoryID
	book.UserID = payload.UserID

	result := r.DB.Table(models.Book{}.TableName()).Save(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	book.User = user
	book.BookCategory = category

	return book, nil
}
