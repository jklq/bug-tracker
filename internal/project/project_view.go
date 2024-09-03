package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/internal/db"
	"github.com/jklq/bug-tracker/internal/helpers"
	"github.com/jklq/bug-tracker/internal/view"
)

func handleProjectView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	project, err := q.GetProjectById(c.Context(), c.Params("projectID"))
	layout := helpers.HtmxLayoutComponent(c)

	if err != nil {
		log.Error(err.Error())

		return view.ErrorView(layout, "Did not find project.").Render(c.Context(), c.Response().BodyWriter())
	}

	tickets, err := q.GetTicketsByProjectId(c.Context(), project.ProjectID)
	if err != nil {
		log.Error(err.Error())

		return view.ErrorView(layout, "Did not find project.").Render(c.Context(), c.Response().BodyWriter())
	}

	projectRole, ok := c.Locals("projectRole").(string)

	if !ok {
		log.Error("c.Locals(\"projectRole\") is not a string")

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return view.ProjectDetailView(view.ProjectDetailViewParams{
		Template:   layout,
		Project:    project,
		Tickets:    tickets,
		SuccessMsg: "",
		Role:       projectRole,
	}).Render(c.Context(), c.Response().BodyWriter())
}
