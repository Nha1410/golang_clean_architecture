package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r UserRepo) FindUserByEmail(email string) (*models.User, error) {
	var user *models.User

	result := r.DB.Table(models.User{}.TableName()).Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
