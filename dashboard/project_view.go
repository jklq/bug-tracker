package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
)

func handleProjectView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	project, err := q.GetProjectById(c.Context(), c.Params("id"))

	if err != nil {
		return c.Render("app/project-details", fiber.Map{"error": "Did not find project."}, helpers.HtmxTemplate(c))
	}

	return c.Render("app/project-details", fiber.Map{"project": project, "noPadding": true}, helpers.HtmxTemplate(c))

}
