package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandlers) DeleteUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID, err :=  strconv.Atoi(ctx.Params("id"))

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code": http.StatusBadRequest,
				"message": "Invalid user ID",
			})
		}

		errDelete := h.userUseCase.DeleteUser(ctx, userID)

		if errDelete != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": "Delete user failed", "errors": errDelete.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(&fiber.Map{"code": http.StatusOK, "message": "Delete user successfully"})
	}
}