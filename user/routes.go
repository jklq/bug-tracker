package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/middleware"
)

func InitModule(router fiber.Router, queries *queryProvider.Queries, db *pgxpool.Pool) {

	onlyUnlogged := router.Group("", middleware.IsNotLoggedIn)

	onlyUnlogged.Get("/login", func(c *fiber.Ctx) error {
		return handleLoginGet(c, queries, db)
	})

	onlyUnlogged.Post("/login", func(c *fiber.Ctx) error {
		return handleLoginPost(c, queries, db)
	})

	onlyUnlogged.Get("/register", func(c *fiber.Ctx) error {
		return handleRegisterGet(c, queries, db)
	})

	onlyUnlogged.Post("/register", func(c *fiber.Ctx) error {
		return handleRegisterPost(c, queries, db)
	})

	protected := router.Group("", middleware.ProtectedRouteMiddleware)

	protected.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})
}
