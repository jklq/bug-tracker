package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
)

func handleTicketDeletion(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	//TODO: CHECK IF USER IS AUTHORIZED TO DELETE PROJECT! ROLE

	err := q.DeleteTicket(c.Context(), c.Params("ticketID"))

	if err != nil {
		print(err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Push-Url", "/app/project/"+c.Params("projectID")+"/view")
	return c.Redirect("/app/project/" + c.Params("projectID") + "/view")
}
