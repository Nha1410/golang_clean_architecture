package repository

import (
	"github.com/team2/real_api/app/models"
)

func (r UserRepo) ListUser() ([]*models.User, error) {
	var users []*models.User

	result := r.DB.Table(models.User{}.TableName()).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
