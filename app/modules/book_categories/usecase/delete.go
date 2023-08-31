package usecase

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func(u BookCategoryUseCase) DeleteBookCategory(ctx *fiber.Ctx, id int) error {
	bookCategory, err := u.bookCategoryRepo.GetByID(id)
	if err != nil {
		return errors.New("Book category not found")
	}

	userID := ctx.Locals("userID").(int)
	if userID != bookCategory.UserID {
		return errors.New("Unauthorized to delete book category")
	}

	err = u.bookCategoryRepo.DeleteBookCategory(bookCategory)
	if err != nil {
		return errors.New("Cannot delete book category")
	}

	return nil
}