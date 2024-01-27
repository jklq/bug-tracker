package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/middleware"
)

func InitModule(router fiber.Router, queries *queryProvider.Queries, db *pgxpool.Pool) {

	protected := router.Group("", middleware.ProtectedRouteMiddleware)

	protected.Get("/", func(c *fiber.Ctx) error {
		return handleProjectListGet(c, queries, db)
	})
	protected.Get("/create", func(c *fiber.Ctx) error {
		return handleProjectCreateView(c, queries, db)
	})
	protected.Post("/create", func(c *fiber.Ctx) error {
		return handleProjectPost(c, queries, db)
	})

	protected.Get("/:projectID/view", func(c *fiber.Ctx) error {
		return handleProjectView(c, queries, db)
	})
	protected.Get("/:projectID/members", func(c *fiber.Ctx) error {
		return handleProjectMemberView(c, queries, db)
	})
	protected.Get("/:projectID/invite", func(c *fiber.Ctx) error {
		return handleProjectMemberInviteSearch(c, queries, db)
	})
	protected.Get("/:projectID/invite/list", func(c *fiber.Ctx) error {
		return handleProjectMemberInvitedList(c, queries, db)
	})
	protected.Post("/:projectID/invite", func(c *fiber.Ctx) error {
		return handleProjectMemberInvite(c, queries, db)
	})
	protected.Delete("/:projectID/uninvite/:userID", func(c *fiber.Ctx) error {
		return handleProjectMemberUninvite(c, queries, db)
	})

	protected.Get("/:projectID/edit", func(c *fiber.Ctx) error {
		return handleEditProjectView(c, queries, db)
	})
	protected.Post("/:projectID/edit", func(c *fiber.Ctx) error {
		return handleEditProjectPost(c, queries, db)
	})
	protected.Post("/:projectID/delete", func(c *fiber.Ctx) error {
		return handleProjectDeletion(c, queries, db)
	})
}
