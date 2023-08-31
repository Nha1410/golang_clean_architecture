package repository

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	"github.com/team2/real_api/config"
	userRepo "github.com/team2/real_api/app/modules/users/repositories"
)

func (r BookCategoryRepo) CreateBookCategory(ctx *fiber.Ctx, payload *models.BookCategoryInput) (*models.BookCategory, error) {
	userID := ctx.Locals("userID").(int)
	user, err := userRepo.NewUserRepo(r.DB).GetUserProfile(userID)
	if err != nil {
			return nil, errors.New("User is not valid")
	}

	var bookCategory = &models.BookCategory{
		Name:  payload.Name,
		Description: payload.Description,
		UserID: userID,
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

	result := r.DB.Table(models.BookCategory{}.TableName()).Create(&bookCategory)
	if result.Error != nil {
		return nil, result.Error
	}

	bookCategory.User = *user
	return bookCategory, nil
}
