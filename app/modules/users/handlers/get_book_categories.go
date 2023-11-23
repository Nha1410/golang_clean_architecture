package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandlers) GetBookCategories() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bookCategoies, err := h.userUseCase.GetBookCategories(ctx)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(bookCategoies)
	}
	
}