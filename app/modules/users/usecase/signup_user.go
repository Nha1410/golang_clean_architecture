package usecase

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/modules/auth"
	"github.com/team2/real_api/app/models"
)

func (u UserUseCase) SignUpUser(ctx *fiber.Ctx, payload *models.SignUpInput) (*models.SignUpResponse, map[string]string) {
	if payload.Password != payload.PasswordConfirmation {
		return nil, map[string]string{"password": "password do not match"}
	}

	// check existing email
	existing := u.userRepo.CheckEmailExists(payload.Email)

	if existing == true {
		return nil, map[string]string{"email": "email existing, please choose another email"}
	}

	createdUser, err := u.userRepo.CreateUser(payload)

	if err != nil {
		return nil, map[string]string{"error": err.Error()}
	}

	token, err := auth.GenerateToken(int(createdUser.ID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError)
	}
 

	return models.ResponseToken(token), nil
}
