package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/auth"
	"github.com/team2/real_api/app/models"
)

func (u UserUseCase) SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.SignUpResponse, error) {
	if payload.Password != payload.PasswordConfirmation {
		return nil, errors.New("passwords do not match")
	}

	// check existing email
	existing := u.userRepo.CheckEmailExists(payload.Email)

	if existing == true {
		return nil, errors.New("email existing, please choose another email")
	}

	createdUser, err := u.userRepo.CreateUser(payload)

	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateToken(int(createdUser.ID))
	fmt.Println(token)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError)
	}
 

	return models.ResponseToken(token), nil
}