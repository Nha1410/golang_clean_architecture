package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandlers) ListUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := h.userUseCase.ListUser(ctx)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(&fiber.Map{"status": http.StatusInternalServerError, "error": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(&fiber.Map{"status": http.StatusOK, "data": users, "error": nil})
	}
}
