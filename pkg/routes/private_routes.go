package routes

import (
	"sample-auth-backend/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/users", controllers.GetUsers)
}
