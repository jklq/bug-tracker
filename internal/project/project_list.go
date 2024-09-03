package project

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/project-tracker/internal/db"
	"github.com/jklq/project-tracker/internal/helpers"
	"github.com/jklq/project-tracker/internal/store"
	"github.com/jklq/project-tracker/internal/view"
)

func handleProjectListGet(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	sess, err := store.Store.Get(c)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userID, ok := sess.Get("user_id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	projects, err := q.GetProjectsByUserIdWithTicketAndMemberInfo(c.Context(), pgtype.Text{String: userID, Valid: true})

	if err != nil {
		log.Println(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	layout := helpers.HtmxLayoutComponent(c)
	return view.ProjectList(layout, projects, "").Render(c.Context(), c.Response().BodyWriter())
}
