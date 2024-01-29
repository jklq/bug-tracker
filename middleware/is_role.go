package middleware

import (
	"github.com/gofiber/fiber/v2"
)

var NotViewer = []string{"project manager", "editor"}

func IsRole(roles ...string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		projectRole, ok := c.Locals("projectRole").(string)
		if !ok {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		for _, role := range roles {
			if projectRole == role {
				return c.Next()
			}
		}

		return c.SendStatus(fiber.StatusForbidden)
	}
}
