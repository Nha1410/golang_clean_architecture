package repository

import (
	"github.com/team2/real_api/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/config"
	"time"
	"errors"
	userRepo "github.com/team2/real_api/app/modules/users/repositories"
)

func (r BookRepo) UpdateBook(ctx *fiber.Ctx,book *models.Book, payload *models.BookInput) (*models.Book, error) {
	user, err := userRepo.NewUserRepo(r.DB).GetUserProfile(int(payload.UserID))
	if err != nil {
			return nil, errors.New("user_id is not valid")
	}
	
	var category models.BookCategory
	err = r.DB.First(&category, payload.BookCategoryID).Error
	if err != nil {
			return nil, errors.New("book_category_id is not valid")
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

	book.Name = payload.Name
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

	book.User = *user
	book.BookCategory = category

	return book, nil
}
