package ticket

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/internal/db"
	"github.com/jklq/bug-tracker/internal/helpers"
	"github.com/jklq/bug-tracker/internal/store"
	"github.com/jklq/bug-tracker/internal/view"
	"github.com/lucsky/cuid"
)

type PostTicketParams struct {
	ProjectID   string
	Title       string `validate:"required,min=3"`
	Description string
	Priority    int16 `validate:"oneof=1 2 3"`
}

// create ticket view within defined project
func handleTicketCreateView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")
	layout := helpers.HtmxLayoutComponent(c)

	return view.CreateTicketView(layout, projectID).Render(c.Context(), c.Response().BodyWriter())
}

// no projectID associated
func handleGeneralTicketCreateView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponent(c)

	userID, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	projects, err := q.GetProjectsByUserId(c.Context(), userID)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return view.GeneralTicketCreateView(layout, projects).Render(c.Context(), c.Response().BodyWriter())
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

	userID, ok := sess.Get("user_id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = q.CreateTicket(c.Context(), queryProvider.CreateTicketParams{
		TicketID:    cuid.New(),
		Title:       params.Title,
		Description: pgtype.Text{String: params.Description, Valid: true},
		Priority:    params.Priority,
		ProjectID:   params.ProjectID,
		Status:      1,
		CreatedBy:   userID,
	})

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tickets, err := q.GetTicketsByProjectId(c.Context(), params.ProjectID)

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	project, err := q.GetProjectById(c.Context(), params.ProjectID)

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", fmt.Sprintf("/app/project/%s/view", params.ProjectID))
	layout := helpers.HtmxLayoutComponent(c)
	return view.ProjectDetailView(view.ProjectDetailViewParams{
		Template:   layout,
		Project:    project,
		Tickets:    tickets,
		SuccessMsg: "Ticket added successfully!"},
	).Render(c.Context(), c.Response().BodyWriter())
}
