package handlers

import (
	user "github.com/team2/real_api/app/modules/users/usecase"
)

type UserHandlers struct {
	userUseCase user.UseCase
}

func NewUserHandlers(userUseCase user.UseCase) *UserHandlers {
	return &UserHandlers{userUseCase: userUseCase}
}
