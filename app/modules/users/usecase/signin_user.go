package usecase

import (
	"errors"
	"github.com/team2/real_api/app/auth"
	"golang.org/x/crypto/bcrypt"
)


func (u UserUseCase) SignInUser(email, password string) (string, error) { 
	user , err := u.userRepo.FindUserByEmail(email)
	if err != nil { 
		return "", errors.New("Unauthorized request")
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte((user.Password)), []byte(password)); err != nil {
		return "", errors.New("Unauthorized request")
	}

	token, err := auth.GenerateToken(int(user.ID)) 
	if err != nil {
        return "", errors.New("Unauthorized request")
    }

	return token, nil
}
