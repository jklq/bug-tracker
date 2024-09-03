package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/project-tracker/internal/db"
)

type DeleteProjectParams struct {
	Id string `validate:"required,min=3"`
}

func handleProjectDeletion(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	//TODO: CHECK IF USER IS AUTHORIZED TO DELETE PROJECT! ROLE

	err := q.DeleteProject(c.Context(), c.Params("projectID"))

	if err != nil {
		log.Error(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/project")
	return c.Redirect("/app/project")
}
