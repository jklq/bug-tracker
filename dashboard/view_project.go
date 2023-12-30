package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
)

func handleProjectView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	project, err := q.GetProjectById(c.Context(), c.Params("id"))

	if err != nil {
		log.Error(err.Error())
		return c.Render("app/project-view", fiber.Map{"error": "Did not find project."}, helpers.HtmxTemplate(c))
	}

	tickets, err := q.GetTicketsByProjectId(c.Context(), project.ProjectID)

	if err != nil {
		log.Error(err.Error())
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding tickets."}, helpers.HtmxTemplate(c))
	}

	return c.Render("app/project-view", fiber.Map{"project": project, "tickets": tickets}, helpers.HtmxTemplate(c))

}
