package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/models"
)

func (h *UserHandlers) SignInUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		payload := models.SignInInput{}
		
		if err := ctx.BodyParser(&payload); err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		fmt.Println(payload.Email)
		token, err := h.userUseCase.Authenticate(payload.Email , payload.Password)

		if err != nil {
			return err
		}
	
		return ctx.JSON(fiber.Map{
			"token": token,
		})

	}
}