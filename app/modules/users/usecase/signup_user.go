package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (u UserUseCase) SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.SignUpResponse, error) {
	// if payload.Password != payload.PasswordConfirmation {
	// 	return nil, errors.New("passwords do not match")
	// }

	// check existing email
	// existing := u.userRepo.CheckEmailExisting(payload.Email)

	// if existing == true {
	// 	return nil, errors.New("email existing, please choose another email.")
	// }

	createdUser, err := u.userRepo.CreateUser(payload)

	if err != nil {
		return nil, err
	}

	return models.ResponseToken(createdUser), nil
}