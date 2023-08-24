package server

import (
	handlerUser "github.com/team2/real_api/app/modules/users/handlers"
	repositoryUser "github.com/team2/real_api/app/modules/users/repositories"
	userUseCase "github.com/team2/real_api/app/modules/users/usecase"
)

func SetupRoutes(server *Server) {

	userRepo := repositoryUser.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := handlerUser.NewUserHandlers(userUseCase)

	api := server.Fiber.Group("/api/v1")

	user := api.Group("/user")
	user.Get("/profile", userHandler.UserProfile())
	user.Post("/signup", userHandler.SignUpUser())
}
