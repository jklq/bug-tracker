package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/store"
	"github.com/jklq/bug-tracker/view"
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

	projects, err := q.GetProjectsByUserIdWithTicketAndMemberInfo(c.Context(), userId)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	layout := helpers.HtmxLayoutComponent(c)
	return view.ProjectList(layout, projects, "").Render(c.Context(), c.Response().BodyWriter())
}
