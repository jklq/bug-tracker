package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
)

func IsProjectMemberMiddleware(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	projectID := c.Params("projectID")
	userID, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	relation, err := q.GetProjectMemberRelation(c.Context(), queryProvider.GetProjectMemberRelationParams{
		ProjectID: projectID,
		UserID:    userID,
	})

	if err != nil {
		log.Warn(err.Error())

		return c.SendStatus(fiber.StatusForbidden)
	}

	c.Locals("projectRole", relation.Role)

	return c.Next()
}
