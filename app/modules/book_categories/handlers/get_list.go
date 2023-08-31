package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (bc BookCategoryHandlers) GetList() fiber.Handler{
	return func(ctx *fiber.Ctx) error {
		bookCategories, err := bc.BookCategoryUseCase.GetList()
		
		if err != nil { 
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(&fiber.Map{"code": http.StatusNotFound, "message": err.Error()})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(bookCategories)
	}
}