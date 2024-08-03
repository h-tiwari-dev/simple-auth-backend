package controllers

import (
	"sample-auth-backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
