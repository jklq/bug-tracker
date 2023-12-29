package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/store"
	"github.com/lucsky/cuid"
)

type PostTicketParams struct {
	Title       string `validate:"required,min=3"`
	Description string
	Priority    int16 `validate:"oneof=1 2 3"`
	AssignedTo  string
}

func handleTicketCreateView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return c.Render("app/create-ticket", fiber.Map{}, helpers.HtmxTemplate(c))

}

func handleTicketCreation(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params PostTicketParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).Render("app/create-ticket", fiber.Map{"error": "Invalid request format"}, helpers.HtmxTemplate(c))
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
	}

	sess, err := store.Store.Get(c)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = q.CreateTicket(c.Context(), queryProvider.CreateTicketParams{
		TicketID:    cuid.New(),
		Title:       params.Title,
		Description: pgtype.Text{String: params.Description, Valid: true},
		Priority:    params.Priority,
		AssignedTo:  pgtype.Text{String: params.AssignedTo, Valid: true},
		ProjectID:   c.Params("projectId"),
		CreatedBy:   userId,
	})

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tickets, err := q.GetTicketsByProjectId(c.Context(), c.Params("projectId"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/tickets")
	return c.Status(fiber.StatusOK).Render("app/tickets", fiber.Map{"tickets": tickets, "success": "Ticket added successfully!"}, helpers.HtmxTemplate(c))
}
