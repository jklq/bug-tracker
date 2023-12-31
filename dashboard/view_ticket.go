package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

func handleTicketDropdownView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, helpers.HtmxTemplate(c))
	}

	return c.Render("app/modules/dropdown", fiber.Map{"ticket": ticket, "action": c.Params("action")}, helpers.HtmxTemplate(c))

}
