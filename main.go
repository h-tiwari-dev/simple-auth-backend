package main

import (
	"log"
	"sample-auth-backend/pkg/configs"
	"sample-auth-backend/pkg/middleware"
	"sample-auth-backend/pkg/routes"
	"sample-auth-backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PrivateRoutes(app)
	routes.PublicRoute(app)
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": false,
			"msg":   "pong",
		})
	})
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}
