package dashboard

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
)

func handleTicketSetStatus(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	status, err := strconv.Atoi(c.Params("status"))

	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Input")
	}

	if status != 1 && status != 2 && status != 0 {
		log.Errorf("status not valid number: %v", status)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Input")
	}

	_, err = q.SetTicketStatus(c.Context(), queryProvider.SetTicketStatusParams{TicketID: c.Params("ticketID"), Status: int16(status)})

	if err != nil {
		log.Errorf(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Redirect("/app/ticket/" + c.Params("ticketID") + "/dropdown/close")
}
