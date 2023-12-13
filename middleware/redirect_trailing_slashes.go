package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RedirectTrailingSlash(c *fiber.Ctx) error {
	url := c.OriginalURL()
	if len(url) > 1 && strings.HasSuffix(url, "/") {
		return c.Redirect(strings.TrimRight(url, "/"), fiber.StatusMovedPermanently)
	}
	return c.Next()
}
