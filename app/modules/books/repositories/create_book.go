package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (r BookRepo) CreateBook(ctx *fiber.Ctx, data *models.BookInput) (*models.Book, error) {
	var user models.User
	var category models.BookCategory
	err := r.DB.First(&user, data.UserID).Error
	if err != nil {
			return nil, errors.New("user_id is not valid")
	}
	err = r.DB.First(&category, data.BookCategoryID).Error
	if err != nil {
			return nil, errors.New("book_category_id is not valid")
	}

	parsedTime, err := time.Parse("01/02/2006", data.PublicDate)
  if err != nil {
    return nil, errors.New("Please provide a valid date with format: mm/dd/yyyy")
  }

	var book = &models.Book{
		Name:  data.Name,
		Author: data.Author,
		PublicDate: parsedTime,
		Description: data.Description,
		BookCategoryID: data.BookCategoryID,
		UserID: data.UserID,
	}

	file, errGetFile := ctx.FormFile("image")
	if errGetFile == nil {
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./assets/image/%s", file.Filename))
		if errSaveFile != nil {
			return nil, errSaveFile
		}
		book.Image = fmt.Sprintf("./assets/image/%s", file.Filename)
	}

	result := r.DB.Table(models.Book{}.TableName()).Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	book.User = user
	book.BookCategory = category

	return book, nil
}
