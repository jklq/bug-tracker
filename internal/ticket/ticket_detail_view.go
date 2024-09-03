package ticket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/internal/db"
	"github.com/jklq/bug-tracker/internal/helpers"
	"github.com/jklq/bug-tracker/internal/view"
)

func handleTicketView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	project, err := q.GetProjectById(c.Context(), c.Params("projectID"))

	if err != nil {
		log.Error(err.Error())
		return view.ErrorView(helpers.HtmxLayoutComponent(c), "Error in finding ticket.").Render(c.Context(), c.Response().BodyWriter())
	}

	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		log.Error(err.Error())
		return view.ErrorView(helpers.HtmxLayoutComponent(c), "Error in finding ticket.").Render(c.Context(), c.Response().BodyWriter())
	}

	if ticket.ProjectID != project.ProjectID {
		log.Info("ticket is not part of project")
		return view.ErrorView(helpers.HtmxLayoutComponent(c), "Error in finding ticket.").Render(c.Context(), c.Response().BodyWriter())
	}

	layout := helpers.HtmxLayoutComponent(c)

	return view.TicketDetailView(layout, project.ProjectID, ticket).Render(c.Context(), c.Response().BodyWriter())
}
