package project

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/view"
)

type EditProjectParams struct {
	Name        string `validate:"required,min=3"`
	Description string
}

func handleEditProjectView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	project, err := q.GetProjectById(c.Context(), c.Params("projectID"))
	layout := helpers.HtmxLayoutComponent(c)

	if err != nil {
		return view.ErrorView(layout, "Did not find project.").Render(c.Context(), c.Response().BodyWriter())
	}
	return view.ProjectEditView(layout, project).Render(c.Context(), c.Response().BodyWriter())
}

func handleEditProjectPost(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	var params EditProjectParams
	projectID := c.Params("projectID")

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	err := helpers.Validate.Struct(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(helpers.TranslateError(err, helpers.Translator)[0].Error())
	}

	_, err = q.UpdateProject(c.Context(), queryProvider.UpdateProjectParams{
		ProjectID:   projectID,
		Name:        params.Name,
		Description: pgtype.Text{String: params.Description, Valid: true},
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

	c.Set("HX-Push-Url", fmt.Sprintf("/app/project/%s/view", projectID))
	return c.Redirect(fmt.Sprintf("/app/project/%s/view", projectID))
}
