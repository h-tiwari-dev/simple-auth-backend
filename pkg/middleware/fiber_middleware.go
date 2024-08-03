package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(
			cors.Config{
				AllowOrigins:     "http://localhost:5173",       // Allow your React app's origin
				AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS", // Allowed HTTP methods
				AllowHeaders:     "Content-Type,Authorization",  // Allowed headers
				ExposeHeaders:    "Content-Length",              // Expose headers
				AllowCredentials: true,                          // Allow credentials (cookies, authorization headers, etc.)
				MaxAge:           300,                           // Maximum age for preflight requests
			},
		),
		// Add simple logger.
		logger.New(),
	)
}
