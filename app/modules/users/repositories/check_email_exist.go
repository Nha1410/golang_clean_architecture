package repository

import (
	"fmt"

	"github.com/team2/real_api/app/models"
)

func (r UserRepo) CheckEmailExists(email string) bool {
	var user *models.User

	result := r.DB.Table(models.User{}.TableName()).Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println("err: ", result.Error)

		return false
	}

	return result.RowsAffected > 0
}
