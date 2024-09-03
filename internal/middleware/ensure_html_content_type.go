package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func EnsureHtmlContentType(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return c.Next()
}
