package ticket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/project-tracker/internal/db"
	"github.com/jklq/project-tracker/internal/helpers"
	"github.com/jklq/project-tracker/internal/store"
	"github.com/jklq/project-tracker/internal/view"
)

func handleAssignedTicketList(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponent(c)

	sess, err := store.Store.Get(c)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userID, ok := sess.Get("user_id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tickets, err := q.GetAssignedTickets(c.Context(), pgtype.Text{String: userID, Valid: true})
	if err != nil {
		log.Error(err.Error())

		return view.ErrorView(layout, "Did not find project.").Render(c.Context(), c.Response().BodyWriter())
	}

	return view.AssignedTicketsView(layout, tickets).Render(c.Context(), c.Response().BodyWriter())
}
