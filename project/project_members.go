package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	queryProvider "github.com/jklq/bug-tracker/db"
	"github.com/jklq/bug-tracker/helpers"
	"github.com/jklq/bug-tracker/store"
	"github.com/jklq/bug-tracker/view"
	"github.com/lucsky/cuid"
)

type InviteUserParams struct {
	UserID string
	Role   string
}

func handleProjectMemberView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")

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
		Project:      project,
		Members:      members,
		InvitedUsers: invitedUsers,
	}

	return view.ProjectMemberDetailView(layout, params).Render(c.Context(), c.Response().BodyWriter())
}

func handleProjectMemberInviteSearch(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	projectID := c.Params("projectID")
	query := c.Query("q")

	if len(query) < 2 {
		return c.SendString("Keep typing!")
	}

	dbparams := queryProvider.SearchUserOutsideProjectParams{ProjectID: projectID, Column2: pgtype.Text{query, true}}

	users, err := q.SearchUserOutsideProject(c.Context(), dbparams)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if len(users) == 0 {
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

	projectId := c.Params("projectID")

	if err := c.BodyParser(&params); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.SendString("Invalid request body")
	}

	if _, ok := validRoles[params.Role]; !ok {
		return c.SendString("Invalid request body")
	}

	sess, err := store.Store.Get(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	dbargs := queryProvider.GetProjectInvitationsByUserAndProjectParams{
		RecipientID: params.UserID,
		ProjectID:   projectId,
	}

	existingInvitations, err := q.GetProjectInvitationsByUserAndProject(c.Context(), dbargs)

	if !ok {
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
		ProjectID:    projectId,
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
	projectId := c.Params("projectID")
	userId := c.Params("userID")

	dbparams := queryProvider.DeleteProjectInvitationParams{
		RecipientID: userId,
		ProjectID:   projectId,
	}

	err := q.DeleteProjectInvitation(c.Context(), dbparams)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("Invitation deleted successfully!")
}

func handleProjectMemberInvitationView(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	layout := helpers.HtmxLayoutComponent(c)
	sess, err := store.Store.Get(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
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

	sess, err := store.Store.Get(c)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	userId, ok := sess.Get("user_id").(string)

	if !ok {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	invitation, err := q.AcceptProjectInvitation(c.Context(), projectID)

	if invitation.RecipientID != userId {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	dbparams := queryProvider.AddUserToProjectParams{
		UserID:    userId,
		ProjectID: projectID,
		Role:      invitation.Role,
	}

	err = q.AddUserToProject(c.Context(), dbparams)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("Invitation accepted successfully!")
}

func handleProjectMemberInviteDecline(c *fiber.Ctx, q *queryProvider.Queries, db *pgxpool.Pool) error {
	invitationId := c.Params("invitationID")

	err := q.DeclineProjectInvitation(c.Context(), invitationId)

	if err != nil {
		log.Error(err.Error())

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("Invitation declined successfully!")
}
