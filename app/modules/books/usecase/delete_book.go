package usecase

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func (u BookUseCase) DeleteBook(ctx *fiber.Ctx, bookID int) error {
	book, err := u.bookRepo.GetBookByID(bookID)
	if err != nil {
		return errors.New("Book not found")
	}

	userID := ctx.Locals("userID").(int)
	if userID != book.UserID {
		return errors.New("Unauthorized to delete book")
	}

	err = u.bookRepo.DeleteBook(book)
	if err != nil {
		return errors.New("Cannot delete book")
	}

	return nil
}
