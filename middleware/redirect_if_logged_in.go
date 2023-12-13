package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jklq/bug-tracker/helpers"
)

// Middleware to check if user is logged in
func RedirectIfLoggedIn(c *fiber.Ctx) error {
	// Get session from store
	isLoggedIn := helpers.IsLoggedIn(c)

	if isLoggedIn {
		return c.Redirect("/app")
	}

	// Continue stack
	return c.Next()
}
