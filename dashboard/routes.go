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

	protected.Get("/project/:projectID/ticket/create", func(c *fiber.Ctx) error {
		return handleTicketCreateView(c, queries, db)
	})
	protected.Post("/project/:projectID/ticket/create", func(c *fiber.Ctx) error {
		return handleTicketCreation(c, queries, db)
	})
	protected.Get("/project/:projectID/ticket/:ticketID/view", func(c *fiber.Ctx) error {
		return handleTicketView(c, queries, db)
	})
	protected.Get("/ticket/:ticketID/dropdown/:action", func(c *fiber.Ctx) error {
		return handleTicketDropdownView(c, queries, db)
	})

	protected.Post("/ticket/:ticketID/status/set/:status<int>", func(c *fiber.Ctx) error {
		return handleTicketSetStatus(c, queries, db)
	})
	protected.Post("/ticket/:ticketID/priority/set/:priority<int>", func(c *fiber.Ctx) error {
		return handleTicketSetPriority(c, queries, db)
	})
}
