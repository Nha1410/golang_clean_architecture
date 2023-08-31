package server

import (
	handlerUser "github.com/team2/real_api/app/modules/users/handlers"
	repositoryUser "github.com/team2/real_api/app/modules/users/repositories"
	userUseCase "github.com/team2/real_api/app/modules/users/usecase"

	handlerBook "github.com/team2/real_api/app/modules/books/handlers"
	repositoryBook "github.com/team2/real_api/app/modules/books/repositories"
	bookUseCase "github.com/team2/real_api/app/modules/books/usecase"
)

func SetupRoutes(server *Server) {
	api := server.Fiber.Group("/api/v1", VerifyToken())

	userRepo := repositoryUser.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := handlerUser.NewUserHandlers(userUseCase)

	user := api.Group("/user")
	user.Get("/profile", userHandler.UserProfile())
	user.Post("/signup", userHandler.SignUpUser())
	user.Post("/signin", userHandler.SignInUser())

	bookRepo := repositoryBook.NewBookRepo(server.DB)
	bookUseCase := bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := handlerBook.NewBookHandlers(bookUseCase)

	book := api.Group("/books")
	book.Get("/", bookHandler.GetBooks())
	book.Post("/new", bookHandler.AddBook())
	book.Put("/:id/edit", bookHandler.EditBook())
	book.Delete("/:id", bookHandler.DeleteBook())
	book.Get("/:id", bookHandler.GetBook())
}
