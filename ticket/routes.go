package ticket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/middleware"
)

func InitModule(router fiber.Router, queries *queryProvider.Queries, db *pgxpool.Pool) {

	protected := router.Group("", middleware.ProtectedRouteMiddleware)

	protected.Get("/create", func(c *fiber.Ctx) error {
		return handleTicketCreateView(c, queries, db)
	})
	protected.Post("/create", func(c *fiber.Ctx) error {
		return handleTicketCreation(c, queries, db)
	})
	protected.Get("/:ticketID/view", func(c *fiber.Ctx) error {
		return handleTicketView(c, queries, db)
	})

	protected.Get("/:ticketID/edit", func(c *fiber.Ctx) error {
		return handleEditTicketView(c, queries, db)
	})
	protected.Post("/:ticketID/edit", func(c *fiber.Ctx) error {
		return handleEditTicketPost(c, queries, db)
	})
	protected.Post("/:ticketID/delete", func(c *fiber.Ctx) error {
		return handleTicketDeletion(c, queries, db)
	})

	protected.Get("/:ticketID/status-dropdown/:action", func(c *fiber.Ctx) error {
		return handleTicketStatusDropdownView(c, queries, db)
	})
	protected.Post("/:ticketID/status/set/:status<int>", func(c *fiber.Ctx) error {
		return handleTicketSetStatus(c, queries, db)
	})

	protected.Get("/:ticketID/assignee", func(c *fiber.Ctx) error {
		return handleGetAssignee(c, queries, db)
	})
	protected.Get("/:ticketID/assignment-dropdown/:action", func(c *fiber.Ctx) error {
		return handleTicketAssignmentDropdownView(c, queries, db)
	})
	protected.Get("/:ticketID/assignTo/search", func(c *fiber.Ctx) error {
		return handleTicketAssigneeSearch(c, queries, db)
	})
	protected.Post("/:ticketID/assignTo/user/:userID", func(c *fiber.Ctx) error {
		return handleTicketDropdownAssign(c, queries, db)
	})
	protected.Post("/:ticketID/priority/set/:priority<int>", func(c *fiber.Ctx) error {
		return handleTicketSetPriority(c, queries, db)
	})
}
