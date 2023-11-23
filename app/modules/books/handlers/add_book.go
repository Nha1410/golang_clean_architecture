package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
	validate "github.com/team2/real_api/app/modules/validate"
	"strconv"
)

func (h *BookHandlers) AddBook() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		BookCategoryID, err :=  strconv.Atoi(ctx.FormValue("book_category_id"))
		if err != nil {
			ctx.JSON(&fiber.Map{"code": http.StatusBadRequest, "message": err.Error()})
		}
	
		payload := models.BookInput{
			Name: ctx.FormValue("name"),
			Author: ctx.FormValue("author"),
			PublicDate: ctx.FormValue("public_date"),
			Description: ctx.FormValue("description"),
			BookCategoryID: BookCategoryID,
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

		createdBook, errors := h.bookUseCase.AddBook(ctx, &payload)
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
