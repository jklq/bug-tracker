package dashboard

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
)

func handleTicketView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	project, err := q.GetProjectById(c.Context(), c.Params("projectID"))

	if err != nil {
		log.Error(err.Error())
		return c.Render("app/project-view", fiber.Map{"error": "Did not find project."}, helpers.HtmxTemplate(c))
	}

	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		log.Error(err.Error())
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, helpers.HtmxTemplate(c))
	}

	if ticket.ProjectID != project.ProjectID {
		log.Info("ticket is not part of project")
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, helpers.HtmxTemplate(c))
	}

	return c.Render("app/ticket-view", fiber.Map{"project": project, "ticket": ticket}, helpers.HtmxTemplate(c))
}

func handleTicketStatusDropdownView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, helpers.HtmxTemplate(c))
	}

	return c.Render("app/modules/ticket-status-dropdown", fiber.Map{"ticket": ticket, "satusDropdownState": c.Params("action")}, "")
}

func handleGetAssignee(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, "")
	}

	user, err := q.GetUserById(c.Context(), ticket.AssignedTo.String)

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

	return c.Render("app/modules/ticket-assignment-dropdown", fiber.Map{"ticket": ticket, "assignmentDropdownState": c.Params("action")}, "")
}

func handleTicketAssigneeSearch(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	query := c.Query("q")
	projectID := c.Params("projectID")
	ticketID := c.Params("ticketID")

	if len(query) < 2 {
		return c.SendString("")
	}

	users, err := q.SearchUserByProject(c.Context(), queryProvider.SearchUserByProjectParams{ProjectID: projectID, Column2: pgtype.Text{String: query, Valid: true}})

	if err != nil {
		log.Warn(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Render("app/modules/ticket-assignment-search-results", fiber.Map{"users": users, "ticketID": ticketID}, "")
}
