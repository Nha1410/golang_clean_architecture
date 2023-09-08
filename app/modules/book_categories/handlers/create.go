package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	validate "github.com/team2/real_api/app/modules/validate"
)

func (h *BookCategoryHandlers) CreateBookCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.BookCategoryInput{
			Name: ctx.FormValue("name"),
			Description: ctx.FormValue("description"),
		}

		errorMessages, validationErr := validate.ValidateFields(&payload)
		if validationErr != nil {
			ctx.Status(http.StatusUnprocessableEntity)

			return ctx.JSON(&fiber.Map{
				"code": http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"errors": errorMessages,
			})
		}

		createdBook, errors := h.BookCategoryUseCase.CreateBookCategory(ctx, &payload)
		if errors != nil {
			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{
				"code": http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"error": errors,
			})
		}

		ctx.Status(http.StatusCreated)
		return ctx.JSON(createdBook)
	}
}
