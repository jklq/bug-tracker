package helpers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsHtmxRequest(c *fiber.Ctx) bool {
	return strings.ToLower(c.Get("HX-Request")) == "true"
}
