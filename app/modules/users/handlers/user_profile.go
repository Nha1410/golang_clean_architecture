package handlers

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func (h *UserHandlers) UserProfile() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := h.userUseCase.UserProfile(ctx)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return ctx.JSON(&fiber.Map{"status": http.StatusUnauthorized, "error": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(&fiber.Map{"status": http.StatusOK, "data": users})
	}
}
