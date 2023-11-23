package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *BookHandlers) DeleteBook() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bookID, err :=  strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": http.StatusBadRequest,
				"message": "Invalid book ID",
			})
		}

		errDelete := h.bookUseCase.DeleteBook(ctx, bookID)

		if errDelete != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(&fiber.Map{"code": http.StatusOK, "message": "Delete book successfully"})
	}
}
