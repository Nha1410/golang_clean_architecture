package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *BookCategoryHandlers) DeleteBookCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		categoryID, err :=  strconv.Atoi(ctx.Params("id"))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": http.StatusBadRequest,
				"message": "Invalid Book Category ID",
			})
		}

		errDelete := h.BookCategoryUseCase.DeleteBookCategory(ctx, categoryID)

		if errDelete != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(&fiber.Map{"code": http.StatusOK, "message": "Delete book category successfully"})
	}
}
