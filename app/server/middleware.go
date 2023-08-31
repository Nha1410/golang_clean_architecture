package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team2/real_api/app/modules/auth"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"strings"
	"net/http"
)

func VerifyToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		excludedRoutes := []string{
			"/api/v1/user/signup",
			"/api/v1/user/signin",
		}

		requestedPath := ctx.Path()
		for _, route := range excludedRoutes {
			if strings.HasPrefix(requestedPath, route) {
				return ctx.Next()
			}
		}
		
		tokenString := strings.TrimPrefix(ctx.Get("Authorization"), "Bearer ")
		if tokenString == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": fiber.StatusUnauthorized, "message": "unauthorized request",
			})
		}

		userID, err := auth.VerifyToken(tokenString)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code": fiber.StatusUnauthorized, "message": "unauthorized request",
			})
		}

		ctx.Locals("userID", userID)

		return ctx.Next()
	}
}

func ConfigureStaticFileMiddleware(app *fiber.App) {
	fs := filesystem.New(filesystem.Config{
    Root:  http.Dir("./assets/image"),
  })

	app.Use(fs)
}
