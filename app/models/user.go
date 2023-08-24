package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string `gorm:"type:varchar(255)" json:"first_name"`
	LastName     string `gorm:"type:varchar(255)" json:"last_name"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Phone string `gorm:"type:varchar(255)" json:"phone"`
}

func (User) TableName() string {
	return "users"
}

type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	FirstName      string    `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName      string    `json:"last_name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone      string    `json:"phone" gorm:"type:varchar(100)"`
}

func FilterUserRecord(user *User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		FirstName:      user.FirstName,
		LastName: user.LastName,
		Email:     user.Email,
		Phone: user.Phone,
	}
}
