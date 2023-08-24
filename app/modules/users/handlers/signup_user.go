package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	"github.com/go-playground/validator/v10"
)

func (h *UserHandlers) SignUpUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignUpInput{}
		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		// Validate the payload using go-validator
		validate := validator.New()
		if err := validate.Struct(payload); err != nil {
			ctx.Status(http.StatusUnprocessableEntity)

			// Custom error response format
			errors := []map[string]interface{}{}
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, map[string]interface{}{
					err.Field(): err.Error(),
				})
			}
			return ctx.JSON(&fiber.Map{
				"code":    http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"errors":  errors,
			})
		}

		createdUser, err := h.userUseCase.SignUpUser(ctx, &payload)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(&fiber.Map{"status": http.StatusCreated, "data": createdUser, "error": nil})
	}
}