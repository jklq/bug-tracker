package user

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/middleware"
	"github.com/jklq/bug-tracker/store"
	"github.com/jklq/bug-tracker/view"
)

func InitModule(router fiber.Router, queries *queryProvider.Queries, db *pgxpool.Pool) {
	router.Get("/logout", func(c *fiber.Ctx) error {
		sess, err := store.Store.Get(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		sess.Destroy()
		return c.Redirect("/user/login")
	})

	onlyUnlogged := router.Group("", middleware.RedirectIfLoggedIn)

	onlyUnlogged.Get("/login", adaptor.HTTPHandler(templ.Handler(view.Login(""))))
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
