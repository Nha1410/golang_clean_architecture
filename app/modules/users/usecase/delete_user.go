package usecase

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (u UserUseCase) DeleteUser(ctx *fiber.Ctx, userID int) error {
	// currentUserID := ctx.Locals("userID").(int)

	// fmt.Println(currentUserID)
	// fmt.Println(ctx.Locals("userID").(int))
	// fmt.Println(currentUserID == userID)
	// if currentUserID == userID {
	// 	return errors.New("you cannot delete yourself")
	// }

    user, err := u.userRepo.GetUserProfile(userID)
    if err != nil {
        return errors.New("user not found")
    }



    // Xóa người dùng
    err = u.userRepo.DeleteUser(user)

    if err != nil {
        return errors.New("cannot delete user")
    }

    return nil
}
