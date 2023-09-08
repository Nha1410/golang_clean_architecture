package server

import (
	handlerUser "github.com/team2/real_api/app/modules/users/handlers"
	repositoryUser "github.com/team2/real_api/app/modules/users/repositories"
	userUseCase "github.com/team2/real_api/app/modules/users/usecase"

	handlerBook "github.com/team2/real_api/app/modules/books/handlers"
	repositoryBook "github.com/team2/real_api/app/modules/books/repositories"
	bookUseCase "github.com/team2/real_api/app/modules/books/usecase"

	handlerbookCategory "github.com/team2/real_api/app/modules/book_categories/handlers"
	repositoryBookCategory "github.com/team2/real_api/app/modules/book_categories/repositories"
	useCaseBookCategory "github.com/team2/real_api/app/modules/book_categories/usecase"
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
	user.Delete("/:id", userHandler.DeleteUser())

	bookRepo := repositoryBook.NewBookRepo(server.DB)
	bookUseCase := bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := handlerBook.NewBookHandlers(bookUseCase)

	book := api.Group("/books")
	book.Get("/", bookHandler.GetBooks())
	book.Post("/new", bookHandler.AddBook())
	book.Put("/:id/edit", bookHandler.EditBook())
	book.Delete("/:id", bookHandler.DeleteBook())
	book.Get("/:id", bookHandler.GetBook())

	bookCategoryRepo := repositoryBookCategory.NewBookCategoryRepo(server.DB)
	bookCategoryUseCase := useCaseBookCategory.NewBookCategoryUseCase(bookCategoryRepo)
	bookCategoryHandler := handlerbookCategory.NewBookCategoryHandlers(bookCategoryUseCase)

	bookCategory := api.Group("/book_categories") 
	bookCategory.Get("/", bookCategoryHandler.GetList())
	bookCategory.Get("/:id", bookCategoryHandler.GetByID())
	bookCategory.Delete("/:id", bookCategoryHandler.DeleteBookCategory())
	bookCategory.Post("/new", bookCategoryHandler.CreateBookCategory())
	bookCategory.Put("/:id/edit", bookCategoryHandler.UpdateBookCategory())
}
