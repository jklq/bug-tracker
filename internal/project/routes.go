package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/project-tracker/internal/db"
	"github.com/jklq/project-tracker/internal/middleware"
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

	protected.Get("/invitation", func(c *fiber.Ctx) error {
		return handleProjectMemberInvitationView(c, queries, db)
	})

	protected.Post("/:projectID/invite/accept", func(c *fiber.Ctx) error {
		return handleProjectMemberInviteAccept(c, queries, db)
	})
	protected.Post("/:projectID/invite/decline", func(c *fiber.Ctx) error {
		return handleProjectMemberInviteDecline(c, queries, db)
	})

	// "/app/project/:projectID" routes past this point
	projectMember := protected.Group("/:projectID", func(c *fiber.Ctx) error {
		return middleware.IsProjectMemberMiddleware(c, queries, db)
	})
	{
		projectMember.Get("/view", func(c *fiber.Ctx) error {
			return handleProjectView(c, queries, db)
		})
		projectMember.Get("/members", func(c *fiber.Ctx) error {
			return handleProjectMemberView(c, queries, db)
		})

		projectMember.Get("/edit", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleEditProjectView(c, queries, db)
		})
		projectMember.Post("/edit", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleEditProjectPost(c, queries, db)
		})
		projectMember.Post("/delete", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleProjectDeletion(c, queries, db)
		})

		// invitation routes
		projectMember.Get("/invite", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleProjectMemberInviteSearch(c, queries, db)
		})
		projectMember.Delete("/uninvite/:userID", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleProjectMemberUninvite(c, queries, db)
		})

		projectMember.Get("/invite/list", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleProjectMemberInvitedList(c, queries, db)
		})
		projectMember.Post("/invite", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleProjectMemberInvite(c, queries, db)
		})
		// remove member
		projectMember.Post("/member/remove/:memberID", middleware.IsRole("project manager"), func(c *fiber.Ctx) error {
			return handleProjectMemberRemove(c, queries, db)
		})

	}

}
