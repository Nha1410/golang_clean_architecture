package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	validate "github.com/team2/real_api/app/modules/validate"
	"strconv"
)

func (h *BookCategoryHandlers) UpdateBookCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bookCategoryID, err :=  strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": http.StatusBadRequest,
				"message": "Invalid book category ID",
			})
		}
	
		payload := models.BookCategoryInput{
			Name: ctx.FormValue("name"),
			Description: ctx.FormValue("description"),
		}

		errorMessages, validationErr := validate.ValidateFields(&payload)
		if validationErr != nil {
			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{
				"code":    http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"errors":  errorMessages,
			})
		}

		updatedBookCategory, errors := h.BookCategoryUseCase.UpdateBookCategory(ctx, bookCategoryID, &payload)
		if errors != nil {
			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{
				"code":    http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"error":   errors,
			})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(updatedBookCategory)
	}
}
