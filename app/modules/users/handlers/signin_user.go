package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (h *UserHandlers) SignInUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignInInput{}
		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		errorMessages, validationErr := h.userUseCase.ValidateFields(&payload)
		if validationErr != nil {
			ctx.Status(http.StatusUnprocessableEntity)
			errorsArray := []fiber.Map{}
			for field, message := range errorMessages {
				errorsArray = append(errorsArray, fiber.Map{field: message})
			}
			return ctx.JSON(&fiber.Map{
				"code": http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"errors": errorsArray,
			})
		}

		token, err := h.userUseCase.SignInUser(payload.Email , payload.Password)
		if err != nil {
			return err
		}
	
		return ctx.JSON(fiber.Map{
			"token": token,
		})

	}
}