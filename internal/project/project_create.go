package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/internal/db"
	"github.com/jklq/bug-tracker/internal/helpers"
	"github.com/jklq/bug-tracker/internal/view"
	"github.com/lucsky/cuid"
)

type PostProjectParams struct {
	Name        string `validate:"required,min=3"`
	Description string
}

func handleProjectCreateView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponent(c)
	return view.ProjectCreateView(layout).Render(c.Context(), c.Response().BodyWriter())
}

func handleProjectPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params PostProjectParams

	if err := c.BodyParser(&params); err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		log.Error(err.Error())

		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
	}

	userID, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = q.CreateProjectByUser(c.Context(), queryProvider.CreateProjectByUserParams{
		ProjectID:   cuid.New(),
		Name:        params.Name,
		Description: pgtype.Text{String: params.Description, Valid: true},
		UserID:      userID,
	})

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/project")
	return c.Redirect("/app/project")
}
