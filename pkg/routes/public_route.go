package routes

import (
	"sample-auth-backend/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(a *fiber.App) {
	route := a.Group("/api/v1/auth")
	route.Get("/sign-up", controllers.SignUp)
}
