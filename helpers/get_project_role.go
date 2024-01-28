package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetProjectRole(c *fiber.Ctx) string {
	if fmt.Sprintf("%T", c.Locals("projectRole")) != "string" {
		log.Error("c.Locals(\"projectRole\") is not a string")
		return ""
	}
	return c.Locals("projectRole").(string)
}
