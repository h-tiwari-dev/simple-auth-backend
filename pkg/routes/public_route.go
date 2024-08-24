package routes

import (
	"sample-auth-backend/app/controllers"
	"sample-auth-backend/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(a *fiber.App) {
	route := a.Group("/api/v1/auth")
	route.Post("/sign-up", controllers.SignUp)
	route.Post("/sign-in", controllers.SignIn)
	route.Get("/google-callback", controllers.GoogleCallbackHandler)
	route.Get("/logged-in", middleware.JWTProtected(), controllers.LoggedIn)
}
