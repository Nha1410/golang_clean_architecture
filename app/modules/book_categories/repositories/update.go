package repository

import (
	"github.com/team2/real_api/app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/config"
	"errors"
	userRepo "github.com/team2/real_api/app/modules/users/repositories"
)

func (r BookCategoryRepo) UpdateBookCategory(ctx *fiber.Ctx,bookCategory *models.BookCategory, payload *models.BookCategoryInput) (*models.BookCategory, error) {
	userID := ctx.Locals("userID").(int)
	user, err := userRepo.NewUserRepo(r.DB).GetUserProfile(userID)
	if err != nil {
			return nil, errors.New("User is not valid")
	}

	if userID != bookCategory.UserID {
		return nil, errors.New("Unauthorized to delete book category")
	}
	
	file, errGetFile := ctx.FormFile("image")
	if errGetFile == nil {
		conf := config.LoadConfig()
		errSaveFile := ctx.SaveFile(file, conf.HTTP.AssetsFolder + file.Filename)
		if errSaveFile != nil {
			return nil, errSaveFile
		}
		bookCategory.Image = conf.HTTP.ImagePath + file.Filename
	}

	bookCategory.Name = payload.Name
	bookCategory.Description = payload.Description
	bookCategory.UserID = userID

	result := r.DB.Table(models.BookCategory{}.TableName()).Save(&bookCategory)
	if result.Error != nil {
		return nil, result.Error
	}

	bookCategory.User = *user
	return bookCategory, nil
}
