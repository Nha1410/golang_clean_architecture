package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func (h *BookHandlers) GetBooks() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		books, err := h.bookUseCase.GetBooks()

		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(books)
	}
}
