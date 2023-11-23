package repository

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	"github.com/team2/real_api/config"
	userRepo "github.com/team2/real_api/app/modules/users/repositories"
)

func (r BookRepo) CreateBook(ctx *fiber.Ctx, payload *models.BookInput) (*models.Book, error) {
	userID := ctx.Locals("userID").(int)
	user, err := userRepo.NewUserRepo(r.DB).GetUserProfile(userID)
	if err != nil {
			return nil, errors.New("User is not valid")
	}

	var category models.BookCategory
	err = r.DB.First(&category, payload.BookCategoryID).Error
	if err != nil {
			return nil, errors.New("book_category_id is not valid")
	}

	parsedTime, err := time.Parse("01/02/2006", payload.PublicDate)
  if err != nil {
    return nil, errors.New("Please provide a valid date with format: mm/dd/yyyy")
  }

	var book = &models.Book{
		Name:  payload.Name,
		Author: payload.Author,
		PublicDate: parsedTime,
		Description: payload.Description,
		BookCategoryID: payload.BookCategoryID,
		UserID: userID,
	}

	file, errGetFile := ctx.FormFile("image")
	if errGetFile == nil {
		conf := config.LoadConfig()
		errSaveFile := ctx.SaveFile(file, conf.HTTP.AssetsFolder + file.Filename)
		if errSaveFile != nil {
			return nil, errSaveFile
		}
		book.Image = conf.HTTP.ImagePath + file.Filename
	}

	result := r.DB.Table(models.Book{}.TableName()).Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	book.User = *user
	book.BookCategory = category

	return book, nil
}
