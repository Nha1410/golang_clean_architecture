package repository

import (
	"github.com/team2/real_api/app/models"
	"golang.org/x/crypto/bcrypt"
)

func (r UserRepo) CreateUser(data *models.SignUpInput) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	var user = &models.User{
		LastName:  data.LastName,
		FirstName: data.FirstName,
		Phone:     data.Phone,
		Email:     data.Email,
		Password:  string(hashedPassword),
	}

	result := r.DB.Table(models.User{}.TableName()).Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}


	return user, nil
}
