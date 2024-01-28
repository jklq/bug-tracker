package middleware

import (
	"github.com/gofiber/fiber/v2"
	queryProvider "github.com/jklq/bug-tracker/db"
)

func IsRole(role string, q *queryProvider.Queries) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if c.Locals("projectRole") != role {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return c.Next()
	}
}
