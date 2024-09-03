package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jklq/project-tracker/internal/store"
)

func GetSession(c *fiber.Ctx) (string, error) {
	sess, err := store.Store.Get(c)

	if err != nil {
		log.Error(err.Error())

		return "", err
	}

	userID, ok := sess.Get("user_id").(string)

	if !ok {
		return "", err
	}

	return userID, nil
}
