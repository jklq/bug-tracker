package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/store"
)

type DeleteProjectParams struct {
	Id string `validate:"required,min=3"`
}

func handleProjectDeletion(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	var params DeleteProjectParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
	}

	//TODO: CHECK IF USER IS AUTHORIZED TO DELETE PROJECT! ROLE

	sess, err := store.Store.Get(c)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		print(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = q.DeleteProject(c.Context(), params.Id)

	if err != nil {
		print(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	projects, err := q.GetProjectsByUserId(c.Context(), userId)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/projects")
	return c.Status(fiber.StatusOK).Render("app/projects", fiber.Map{"projects": projects, "success": "Project deleted successfully!"}, helpers.HtmxTemplate(c))
}
