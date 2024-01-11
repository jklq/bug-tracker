package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

func handleProjectCreateView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	return c.Render("app/create-project", fiber.Map{}, helpers.HtmxTemplate(c))

}

func handleProjectPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params PostProjectParams

	if err := c.BodyParser(&params); err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).Render("app/create-project", fiber.Map{"error": "Invalid request format"}, helpers.HtmxTemplate(c))
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		log.Error(err.Error())

		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
	}

	sess, err := store.Store.Get(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	_, err = q.CreateProjectByUser(c.Context(), queryProvider.CreateProjectByUserParams{
		ProjectID:   cuid.New(),
		Name:        helpers.CleanHTML(params.Name),
		Description: pgtype.Text{String: helpers.CleanHTML(params.Description), Valid: true},
		UserID:      userId,
	})

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	projects, err := q.GetProjectsByUserId(c.Context(), userId)

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/projects")
	return c.Status(fiber.StatusOK).Render("app/projects", fiber.Map{"projects": projects, "success": "Project added successfully!"}, helpers.HtmxTemplate(c))
}
