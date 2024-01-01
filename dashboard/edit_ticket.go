package dashboard

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
)

type EditTicketParams struct {
	Title       string `validate:"required,min=3"`
	Description string
	Status      string `validate:"required,oneof=0 1 2,number"`
	Priority    string `validate:"required,oneof=1 2 3,number"`
}

func handleEditTicketView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	project, err := q.GetProjectById(c.Context(), c.Params("projectID"))

	if err != nil {
		return c.Render("app/edit-ticket", fiber.Map{"error": "Did not find project."}, helpers.HtmxTemplate(c))
	}

	ticket, err := q.GetTicketById(c.Context(), c.Params("ticketID"))

	if err != nil {
		return c.Render("app/edit-ticket", fiber.Map{"error": "Did not find ticket."}, helpers.HtmxTemplate(c))
	}

	return c.Render("app/edit-ticket", fiber.Map{"project": project, "ticket": ticket}, helpers.HtmxTemplate(c))
}

func handleEditTicketPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params EditTicketParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
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
		TicketID:    c.Params("TicketID"),
		Title:       params.Title,
		Description: pgtype.Text{String: params.Description, Valid: true},
		Status:      int16(status),
		Priority:    int16(priority),
	})

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//TODO: CHECK IF USER IS AUTHORIZED TO UPDATE PROJECT! ROLE

	// sess, err := store.Store.Get(c)

	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	// userId, ok := sess.Get("user_id").(string)

	// if !ok {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	c.Set("HX-Push-Url", "/app/project/"+c.Params("projectID")+"/ticket/"+c.Params("ticketID")+"/view")

	return c.Redirect("/app/project/" + c.Params("projectID") + "/ticket/" + c.Params("ticketID") + "/view")
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

	return c.Redirect("/app/ticket/" + c.Params("ticketID") + "/dropdown/close")
}

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

	_, err = q.SetTicketPrioirty(c.Context(), queryProvider.SetTicketPrioirtyParams{TicketID: c.Params("ticketID"), Priority: int16(priority)})

	if err != nil {
		log.Errorf(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render("app/modules/ticket-priority-col", fiber.Map{"ticket": fiber.Map{"Priority": priority, "TicketID": c.Params("ticketID")}}, "")
}
