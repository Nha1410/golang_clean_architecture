package repository

import "github.com/team2/real_api/app/models"

func (u *UserRepo) FindUserByEmail(email string) (*models.User, error) {
	var user *models.User

	result := u.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
        return nil, result.Error
    }
    return result.RowsAffected, nil
}