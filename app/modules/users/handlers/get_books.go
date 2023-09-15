package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandlers) GetBooks() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := h.userUseCase.GetBooks(ctx)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(users)
	}
}
