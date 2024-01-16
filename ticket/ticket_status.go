package ticket

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/view"
	"github.com/mitchellh/mapstructure"
)

func handleTicketStatusDropdownView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return c.Render("app/project-view", fiber.Map{"error": "Error in finding ticket."}, helpers.HtmxTemplate(c))
	}

	var ticketListTicket view.TicketList_ticket
	mapstructure.Decode(ticket, &ticketListTicket)

	return view.TicketStatusDropdown(ticketListTicket, c.Params("action")).Render(c.Context(), c.Response().BodyWriter())
}

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

	return c.Redirect(fmt.Sprintf("/app/project/%s/ticket/%s/status-dropdown/close", c.Params("projectID"), c.Params("ticketID")))
}
