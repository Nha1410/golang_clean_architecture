package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r UserRepo) GetUserProfile(userID int) (*models.User, error) {
	var user *models.User

	result := r.DB.Table(models.User{}.TableName()).Where("id = ?", userID).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
