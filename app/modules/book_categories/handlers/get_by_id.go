package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *BookCategoryHandlers) GetByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bookCategoryID, err :=  strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": http.StatusBadRequest,
				"message": "Invalid book category ID",
			})
		}

		book, err := h.BookCategoryUseCase.GetBookCategoryByID(bookCategoryID)

		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(book)
	}
}
