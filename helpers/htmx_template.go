package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func HtmxTemplate(c *fiber.Ctx) string {
	if IsHtmxRequest(c) {
		return ""
	}
	return "layouts/main"
}
