package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (u UserUseCase) DeleteUser(ctx *fiber.Ctx, userID int) error {
    user, err := u.userRepo.GetUserProfile(userID)
    if err != nil {
        return errors.New("user not found")
    }

    err = u.userRepo.DeleteUser(user)

    if err != nil {
        return errors.New("cannot delete user")
    }

    return nil
}
