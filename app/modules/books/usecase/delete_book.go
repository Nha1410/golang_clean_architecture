package usecase

import (
	"errors"
)

func (u BookUseCase) DeleteBook(bookID int) error {
	err := u.bookRepo.DeleteBook(bookID)

	if err != nil {
		return errors.New("Cannot delete book")
	}

	return nil
}
