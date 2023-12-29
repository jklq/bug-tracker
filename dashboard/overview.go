package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/store"
)

func handleProjectListGet(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	sess, err := store.Store.Get(c)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	projects, err := q.GetProjectsByUserId(c.Context(), userId)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render("app/projects", fiber.Map{"projects": projects, "url": c.Path()}, helpers.HtmxTemplate(c))
}

func handleDashboardPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return nil
}
