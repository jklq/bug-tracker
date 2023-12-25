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

func handleProjectCreateView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return c.Render("app/create-project", fiber.Map{}, helpers.HtmxTemplate(c))

}

func handleProjectPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params PostProjectParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).Render("app/create-project", fiber.Map{"error": "Invalid request format"}, helpers.HtmxTemplate(c))
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Render("app/create-project", fiber.Map{"error": helpers.TranslateError(err, helpers.Translator)[0].Error()}, helpers.HtmxTemplate(c))
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

	projects, err := q.GetProjectsByUserId(c.Context(), userId)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/projects")
	// TODO: this does not work
	return c.Status(fiber.StatusOK).Render("app/projects", fiber.Map{"projects": projects, "success": "Project added successfully!"}, helpers.HtmxTemplate(c))
}
