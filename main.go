package main

import (
	"sample-auth-backend/pkg/configs"
	"sample-auth-backend/pkg/middleware"
	"sample-auth-backend/pkg/routes"
	"sample-auth-backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)
	routes.SwaggerRoute(app)

	utils.StartServer(app)
}
