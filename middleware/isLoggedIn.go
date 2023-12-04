package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jklq/bug-tracker/store"
)

// Middleware to check if user is logged in
func IsNotLoggedIn(c *fiber.Ctx) error {
	// Get session from store
	sess, err := store.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	print(sess.Get("user_id") == nil)
	print(sess.Get("user_id"))

	if sess.Get("user_id") != nil {
		return c.Redirect("/")
	}

	// Continue stack
	return c.Next()
}
