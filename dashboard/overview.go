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

type PostProjectParams struct {
	Name        string `validate:"required,min=3"`
	Description string
}

func handleDashboardGet(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	sess, err := store.Store.Get(c)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	projects, err := q.GetProjectsByUserId(c.Context(), userId)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render("app/dashboard", fiber.Map{"projects": projects})
}

func handleProjectPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params PostProjectParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).Render("dashboard", fiber.Map{"error": "Invalid request format"})
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Render("dashboard", fiber.Map{"error": helpers.TranslateError(err, helpers.Translator)[0].Error()})
	}

	sess, err := store.Store.Get(c)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = q.CreateProjectByUser(c.Context(), queryProvider.CreateProjectByUserParams{
		ProjectID:   cuid.New(),
		Name:        params.Name,
		Description: pgtype.Text{String: params.Description, Valid: true},
		UserID:      userId,
	})

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Redirect("/app")
}

func handleDashboardPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return nil
}
