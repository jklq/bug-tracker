package ticket

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/project-tracker/internal/db"
	"github.com/jklq/project-tracker/internal/view"
)

func handleTicketSetPriority(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	priority, err := strconv.Atoi(c.Params("priority"))

	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Input")
	}

	if priority != 1 && priority != 2 && priority != 3 {
		log.Errorf("priority not valid number: %v", priority)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Input")
	}

	ticket, err := q.SetTicketPrioirty(c.Context(), queryProvider.SetTicketPrioirtyParams{TicketID: c.Params("ticketID"), Priority: int16(priority)})

	if err != nil {
		log.Errorf(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	respParams := view.TicketPriorityTicker_ticket{TicketID: ticket.TicketID, ProjectID: ticket.ProjectID, Priority: ticket.Priority}

	return view.TicketPriorityTicker(respParams).Render(c.Context(), c.Response().BodyWriter())
}
