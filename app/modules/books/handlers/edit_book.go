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

		BookCategoryID, err :=  strconv.Atoi(ctx.FormValue("book_category_id"))
		if err != nil {
			ctx.JSON(&fiber.Map{"code": http.StatusBadRequest, "message": err.Error()})
		}

		UserID, err :=  strconv.Atoi(ctx.FormValue("user_id"))
		if err != nil {
			ctx.JSON(&fiber.Map{"code": http.StatusBadRequest, "message": err.Error()})
		}
	
		payload := models.BookInput{
			Name: ctx.FormValue("name"),
			Author: ctx.FormValue("author"),
			PublicDate: ctx.FormValue("public_date"),
			Description: ctx.FormValue("description"),
			BookCategoryID: uint(BookCategoryID),
			UserID: uint(UserID),
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
