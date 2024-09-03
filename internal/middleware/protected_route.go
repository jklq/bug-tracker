package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jklq/bug-tracker/internal/store"
)

// Middleware to check if user is logged in
func ProtectedRouteMiddleware(c *fiber.Ctx) error {
	// Get session from store
	sess, err := store.Store.Get(c)
	if err != nil {
		print(err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if sess.Get("user_id") == nil {
		return c.Redirect("/user/login")
	}

	// Continue stack
	return c.Next()
}
