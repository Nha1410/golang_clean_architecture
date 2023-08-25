package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (h *UserHandlers) SignUpUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignUpInput{}
		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"code": http.StatusBadRequest, "message": err.Error()})
		}

		errorMessages, validationErr := h.userUseCase.ValidateFields(&payload)
		if validationErr != nil {
			ctx.Status(http.StatusUnprocessableEntity)

			return ctx.JSON(&fiber.Map{
				"code": http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"errors": errorMessages,
			})
		}

		createdUser, errors := h.userUseCase.SignUpUser(ctx, &payload)
		if errors != nil {
			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{
				"code": http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"error": errors,
			})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(createdUser)
	}
}
