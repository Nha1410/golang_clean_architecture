package usecase

import (
	"github.com/team2/real_api/app/auth"
	"golang.org/x/crypto/bcrypt"
)


func (u UserUseCase) Authenticate(email, password string) (string, error) { 
	
	user , err := u.userRepo.FindUserByEmail(email)

	if err != nil { 
		return "", err
	}

	// Verify the password


	if err := bcrypt.CompareHashAndPassword([]byte((user.Password)), []byte(password)); err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(int(user.ID)) 
	if err != nil {
        return "", err
    }

	return token, nil
}