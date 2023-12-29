package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/middleware"
)

func InitModule(router fiber.Router, queries *queryProvider.Queries, db *pgxpool.Pool) {

	protected := router.Group("", middleware.ProtectedRouteMiddleware)

	protected.Get("/", func(c *fiber.Ctx) error {
		return c.Render("app/overview", nil)
	})
	protected.Get("/projects", func(c *fiber.Ctx) error {
		return handleProjectListGet(c, queries, db)
	})
	protected.Get("/project/create", func(c *fiber.Ctx) error {
		return handleProjectCreateView(c, queries, db)
	})
	protected.Post("/project/create", func(c *fiber.Ctx) error {
		return handleProjectPost(c, queries, db)
	})
	protected.Get("/project/:id/view", func(c *fiber.Ctx) error {
		return handleProjectView(c, queries, db)
	})
	protected.Get("/project/:id/edit", func(c *fiber.Ctx) error {
		return handleEditProjectView(c, queries, db)
	})
	protected.Post("/project/:id/edit", func(c *fiber.Ctx) error {
		return handleEditProjectPost(c, queries, db)
	})
	protected.Post("/project/:id/delete", func(c *fiber.Ctx) error {
		return handleProjectDeletion(c, queries, db)
	})

	protected.Post("/project/:projectId/ticket/create", func(c *fiber.Ctx) error {
		return handleTicketCreation(c, queries, db)
	})

}
