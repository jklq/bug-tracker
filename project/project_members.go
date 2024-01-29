package project

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/view"
	"github.com/lucsky/cuid"
)

type InviteUserParams struct {
	UserID string
	Role   string
}

func handleProjectMemberView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")

	userID, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	project, err := q.GetProjectById(c.Context(), projectID)
	layout := helpers.HtmxLayoutComponent(c)

	if err != nil {
		log.Error(err.Error())

		return view.ErrorView(layout, "Did not find project.").Render(c.Context(), c.Response().BodyWriter())
	}

	members, err := q.GetProjectMembersWithRoles(c.Context(), project.ProjectID)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	invitedUsers, err := q.GetUsersWithOpenProjectInvitations(c.Context(), projectID)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	params := view.ProjectMemberDetailViewParams{
		UserID:       userID,
		Project:      project,
		Members:      members,
		InvitedUsers: invitedUsers,
	}

	projectRole := helpers.GetProjectRole(c)

	return view.ProjectMemberDetailView(layout, params, projectRole).Render(c.Context(), c.Response().BodyWriter())
}

func handleProjectMemberInviteSearch(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")
	query := c.Query("q")

	if len(query) < 2 {
		return c.SendString("Keep typing!")
	}

	dbparams := queryProvider.SearchUserOutsideProjectParams{ProjectID: projectID, Column2: pgtype.Text{String: query, Valid: true}}

	users, err := q.SearchUserOutsideProject(c.Context(), dbparams)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(users) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.SendString("No users found!")
	}

	return view.InviteUserSearchResultView(users[0]).Render(c.Context(), c.Response().BodyWriter())
}

func handleProjectMemberInvitedList(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")

	project, err := q.GetProjectById(c.Context(), projectID)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	invitedUsers, err := q.GetUsersWithOpenProjectInvitations(c.Context(), projectID)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return view.InvitedUserList(invitedUsers, project).Render(c.Context(), c.Response().BodyWriter())
}

var validRoles = map[string]bool{
	"viewer":          true,
	"project manager": true,
	"editor":          true,
}

func handleProjectMemberInvite(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	var params InviteUserParams

	projectID := c.Params("projectID")

	if err := c.BodyParser(&params); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("Invalid request body")
	}

	if _, ok := validRoles[params.Role]; !ok {
		return c.SendString("Invalid request body")
	}

	userId, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	dbargs := queryProvider.GetProjectInvitationsByUserAndProjectParams{
		RecipientID: params.UserID,
		ProjectID:   projectID,
	}

	existingInvitations, err := q.GetProjectInvitationsByUserAndProject(c.Context(), dbargs)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(existingInvitations) > 0 {
		c.Status(fiber.StatusConflict)
		return c.SendString("Invitation already sent!")
	}

	dbparams := queryProvider.CreateProjectInvitationParams{
		InvitationID: cuid.New(),
		RecipientID:  params.UserID,
		SenderID:     pgtype.Text{String: userId, Valid: true},
		Role:         params.Role,
		ProjectID:    projectID,
	}

	err = q.CreateProjectInvitation(c.Context(), dbparams)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("HX-Trigger", "invite-user")

	return c.SendString("Invitation sent successfully!")
}

func handleProjectMemberUninvite(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")
	userId := c.Params("userID")

	invitations, err := q.GetProjectInvitationsByUserAndProject(c.Context(), queryProvider.GetProjectInvitationsByUserAndProjectParams{
		RecipientID: userId,
		ProjectID:   projectID,
	})

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(invitations) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.SendString("No invitation found!")
	}

	invitation := invitations[0]

	if invitation.RecipientID != userId {
		return c.SendStatus(fiber.StatusForbidden)
	}

	dbparams := invitation.InvitationID

	err = q.DeleteProjectInvitation(c.Context(), dbparams)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("Invitation deleted successfully!")
}

func handleProjectMemberInvitationView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponent(c)

	userId, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	invitations, err := q.GetProjectInvitationsByUser(c.Context(), userId)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return view.ProjectMemberInvitationView(layout, invitations).Render(c.Context(), c.Response().BodyWriter())
}

func handleProjectMemberInviteAccept(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {

	projectID := c.Params("projectID")

	userId, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tx, err := db.Begin(c.Context())

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer tx.Rollback(c.Context())
	qtx := q.WithTx(tx)

	invitations, err := qtx.GetProjectInvitationsByUserAndProject(c.Context(), queryProvider.GetProjectInvitationsByUserAndProjectParams{
		RecipientID: userId,
		ProjectID:   projectID,
	})

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(invitations) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.SendString("No invitation found!")
	}

	invitation := invitations[0]

	err = qtx.AddUserToProject(c.Context(), queryProvider.AddUserToProjectParams{
		UserID:    userId,
		ProjectID: projectID,
		Role:      invitation.Role,
	})

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = qtx.DeleteProjectInvitation(c.Context(), invitation.InvitationID)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = tx.Commit(c.Context())

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("Invitation accepted successfully!")
}

func handleProjectMemberInviteDecline(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")

	userID, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	invitations, err := q.GetProjectInvitationsByUserAndProject(c.Context(), queryProvider.GetProjectInvitationsByUserAndProjectParams{
		RecipientID: userID,
		ProjectID:   projectID,
	})

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(invitations) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.SendString("No invitation found!")
	}

	invitation := invitations[0]

	if invitation.RecipientID != userID {
		return c.SendStatus(fiber.StatusForbidden)
	}

	fmt.Println("Recipient", invitation)
	fmt.Println("UserID", userID)

	err = q.DeleteProjectInvitation(c.Context(), invitation.InvitationID)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("Invitation declined successfully!")
}

// handleProjectMemberRemove
func handleProjectMemberRemove(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")
	userID := c.Params("userID")

	userId, err := helpers.GetSession(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if userId == userID {
		return c.SendStatus(fiber.StatusForbidden)
	}

	err = q.RemoveUserFromProject(c.Context(), queryProvider.RemoveUserFromProjectParams{
		UserID:    userID,
		ProjectID: projectID,
	})

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("User removed from project successfully!")
}
