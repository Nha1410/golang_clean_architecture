package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	validate "github.com/team2/real_api/app/modules/validate"
	"strconv"
)

func (h *BookHandlers) EditBook() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bookID, err :=  strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": http.StatusBadRequest,
				"message": "Invalid book ID",
			})
		}

		payload := models.BookInput{}
		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"code": http.StatusBadRequest, "message": err.Error()})
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

		updatedBook, errors := h.bookUseCase.EditBook(ctx, bookID, &payload)
		if errors != nil {
			ctx.Status(http.StatusUnprocessableEntity)
			return ctx.JSON(&fiber.Map{
				"code":    http.StatusUnprocessableEntity,
				"message": "Unprocessable Content",
				"error":   errors,
			})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(updatedBook)
	}
}
