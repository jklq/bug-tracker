package ticket

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/view"
)

type EditTicketParams struct {
	Title       string `validate:"required,min=3"`
	Description string
	Status      string `validate:"required,oneof=0 1 2,number"`
	Priority    string `validate:"required,oneof=1 2 3,number"`
}

func handleEditTicketView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponent(c)

	userID, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = q.GetProjectMemberRelation(c.Context(), queryProvider.GetProjectMemberRelationParams{
		ProjectID: c.Params("projectID"),
		UserID:    userID,
	})

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusForbidden)
	}

	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return view.ErrorView(layout, "Did not find ticket.").Render(c.Context(), c.Response().BodyWriter())
	}

	return view.TicketEditView(layout, ticket).Render(c.Context(), c.Response().BodyWriter())
}

func handleEditTicketPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	var params EditTicketParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
	}

	// first get project id from session
	userId, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// then check if user is member of project
	_, err = q.GetProjectMemberRelation(c.Context(), queryProvider.GetProjectMemberRelationParams{
		ProjectID: c.Params("projectID"),
		UserID:    userId,
	})

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusForbidden)
	}

	status, err := strconv.Atoi(params.Status)

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	priority, err := strconv.Atoi(params.Priority)

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err = q.UpdateTicket(c.Context(), queryProvider.UpdateTicketParams{
		TicketID:    c.Params("ticketID"),
		Title:       params.Title,
		Description: pgtype.Text{String: params.Description, Valid: true},
		Status:      int16(status),
		Priority:    int16(priority),
	})

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", fmt.Sprintf("/app/project/%s/ticket/%s/view", c.Params("projectID"), c.Params("ticketID")))
	return c.Redirect(fmt.Sprintf("/app/project/%s/ticket/%s/view", c.Params("projectID"), c.Params("ticketID")))
}
