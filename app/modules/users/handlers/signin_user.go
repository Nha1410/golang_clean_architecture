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
			return ctx.JSON(&fiber.Map{"code": http.StatusBadRequest, "message": err.Error()})
		}

		_, validationErr := h.userUseCase.ValidateFields(&payload)
		if validationErr != nil {
			ctx.Status(http.StatusUnprocessableEntity)

			return ctx.JSON(&fiber.Map{
				"code": http.StatusUnprocessableEntity,
				"message": "Invalid input",
			})
		}

		token, err := h.userUseCase.SignInUser(payload.Email , payload.Password)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(&fiber.Map{"code": http.StatusUnauthorized, "message": err.Error()})
		}
	
		return ctx.JSON(fiber.Map{
			"token": token,
			"code": http.StatusOK,
		})
	}
}
