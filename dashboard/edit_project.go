package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
)

type EditProjectParams struct {
	Id          string `validate:"required,min=3"`
	Name        string `validate:"required,min=3"`
	Description string
}

func handleEditProjectView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	project, err := q.GetProjectById(c.Context(), c.Params("id"))

	if err != nil {
		return c.Render("app/edit-project", fiber.Map{"error": "Did not find project."}, helpers.HtmxTemplate(c))
	}
	return c.Render("app/edit-project", fiber.Map{"project": project}, helpers.HtmxTemplate(c))
}

func handleEditProjectPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params EditProjectParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
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

	project, err := q.UpdateProject(c.Context(), queryProvider.UpdateProjectParams{
		ProjectID:   params.Id,
		Name:        params.Name,
		Description: pgtype.Text{String: params.Description, Valid: true},
	})

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/project/"+params.Id+"/view")

	return c.Status(fiber.StatusOK).Render("app/project-view", fiber.Map{"project": project}, helpers.HtmxTemplate(c))
}
