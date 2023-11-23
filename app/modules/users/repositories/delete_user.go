package repository

import (
	"github.com/team2/real_api/app/models"
)

func(r UserRepo) DeleteUser(user *models.User) error {
	if err := r.DB.Model(&user).Association("Books").Delete(user.Books); err != nil {
		return err
	}

	if err := r.DB.Model(&user).Association("BookCategories").Delete(user.BookCategories); err != nil {
		return err
	}

	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}