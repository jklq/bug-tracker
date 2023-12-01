package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jklq/bug-tracker/db"
)

type Article struct {
	Title    string
	Content  string
	Category string
}

func InitRadar(router fiber.Router, queries *db.Queries) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Add login, register, etc here")
	})
}
