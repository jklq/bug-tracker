package ticket

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/view"
)

func handleGetAssignee(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, "")
	}

	user, err := q.GetUserById(c.Context(), ticket.AssigneeID.String)

	if err != nil {
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding user."}, "")
	}

	return c.SendString(fmt.Sprintf("%s (%s)", user.Username, user.Email))
}

func handleTicketAssignmentDropdownView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, "")
	}

	return view.TicketAssignmentDropdown(ticket, c.Params("action")).Render(c.Context(), c.Response().BodyWriter())
}

func handleTicketAssigneeSearch(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	query := c.Query("q")
	ticketID := c.Params("ticketID")

	if len(query) < 2 {
		return c.SendString("")
	}

	ticket, err := q.GetTicketById(c.Context(), ticketID)

	if err != nil {
		log.Warn(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	users, err := q.SearchUserByProject(c.Context(), queryProvider.SearchUserByProjectParams{ProjectID: ticket.ProjectID, Column2: pgtype.Text{String: query, Valid: true}})

	if err != nil {
		log.Warn(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return view.TicketAssignmentSearchResults(ticketID, users).Render(c.Context(), c.Response().BodyWriter())
}

func handleTicketDropdownAssign(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	ticketID := c.Params("ticketID")
	userID := c.Params("userID")

	_, err := q.SetTicketAssignee(c.Context(), queryProvider.SetTicketAssigneeParams{
		TicketID:   ticketID,
		AssigneeID: pgtype.Text{String: userID, Valid: true},
	})

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	ticket, err := q.GetTicketById(c.Context(), ticketID)

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return view.TicketAssignmentDropdown(ticket, "close").Render(c.Context(), c.Response().BodyWriter())
}
