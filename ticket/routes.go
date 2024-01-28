package ticket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/middleware"
)

func InitModule(router fiber.Router, queries *queryProvider.Queries, db *pgxpool.Pool) {
	protected := router.Group("", middleware.ProtectedRouteMiddleware)

	// General routes here

	protected.Get("/ticket", func(c *fiber.Ctx) error {
		return handleAssignedTicketList(c, queries, db)
	})
	protected.Get("/ticket/new", func(c *fiber.Ctx) error {
		return handleGeneralTicketCreateView(c, queries, db)
	})

	protected.Post("/ticket/create", func(c *fiber.Ctx) error {
		return handleTicketCreation(c, queries, db)
	})

	protected.Get("/ticket/:ticketID/assign/search", func(c *fiber.Ctx) error {
		return handleTicketAssigneeSearch(c, queries, db)
	})
	protected.Post("/ticket/:ticketID/assignTo/user/:userID", func(c *fiber.Ctx) error {
		return handleTicketDropdownAssign(c, queries, db)
	})

	// Spesific routes here
	projectTicket := protected.Group("/project/:projectID/ticket", func(c *fiber.Ctx) error {
		return middleware.IsProjectMemberMiddleware(c, queries, db)
	})
	{
		projectTicket.Get("/create", func(c *fiber.Ctx) error {
			return handleTicketCreateView(c, queries, db)
		})

		projectTicket.Get("/:ticketID/view", func(c *fiber.Ctx) error {
			return handleTicketView(c, queries, db)
		})

		projectTicket.Get("/:ticketID/edit", func(c *fiber.Ctx) error {
			return handleEditTicketView(c, queries, db)
		})
		projectTicket.Post("/:ticketID/edit", func(c *fiber.Ctx) error {
			return handleEditTicketPost(c, queries, db)
		})
		projectTicket.Post("/:ticketID/delete", func(c *fiber.Ctx) error {
			return handleTicketDeletion(c, queries, db)
		})

		projectTicket.Get("/:ticketID/status-dropdown/:action", func(c *fiber.Ctx) error {
			return handleTicketStatusDropdownView(c, queries, db)
		})
		projectTicket.Post("/:ticketID/status/set/:status<int>", func(c *fiber.Ctx) error {
			return handleTicketSetStatus(c, queries, db)
		})

		projectTicket.Get("/:ticketID/assignee", func(c *fiber.Ctx) error {
			return handleGetAssignee(c, queries, db)
		})
		projectTicket.Get("/:ticketID/assignment-dropdown/:action", func(c *fiber.Ctx) error {
			return handleTicketAssignmentDropdownView(c, queries, db)
		})

		projectTicket.Post("/:ticketID/priority/set/:priority<int>", func(c *fiber.Ctx) error {
			return handleTicketSetPriority(c, queries, db)
		})
	}

}
