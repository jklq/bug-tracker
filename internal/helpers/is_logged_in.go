package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jklq/project-tracker/internal/store"
)

// Middleware to check if user is logged in
func IsLoggedIn(c *fiber.Ctx) bool {
	// Get session from store
	sess, err := store.Store.Get(c)

	if err != nil {
		log.Error("Store.Get() Failed. This is critical, and should never happen.")
		return false
	}

	if sess.Get("user_id") != nil {
		return true
	}

	return false
}
