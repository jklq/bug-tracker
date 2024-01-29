// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addUserToProject = `-- name: AddUserToProject :exec
INSERT INTO user_projects (user_id, project_id, role) VALUES ($1, $2, $3)
`

type AddUserToProjectParams struct {
	UserID    string
	ProjectID string
	Role      string
}

func (q *Queries) AddUserToProject(ctx context.Context, arg AddUserToProjectParams) error {
	_, err := q.db.Exec(ctx, addUserToProject, arg.UserID, arg.ProjectID, arg.Role)
	return err
}

const createProject = `-- name: CreateProject :one
INSERT INTO projects (
  project_id,
  name,
  description,
  created_by
) VALUES ($1, $2, $3, $4)
RETURNING project_id, name, description, created_by, created_at, updated_at
`

type CreateProjectParams struct {
	ProjectID   string
	Name        string
	Description pgtype.Text
	CreatedBy   pgtype.Text
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, createProject,
		arg.ProjectID,
		arg.Name,
		arg.Description,
		arg.CreatedBy,
	)
	var i Project
	err := row.Scan(
		&i.ProjectID,
		&i.Name,
		&i.Description,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createProjectByUser = `-- name: CreateProjectByUser :one
WITH new_project AS (
    INSERT INTO projects (project_id, name, description, created_by)
    VALUES ($1, $2, $3, $4)
    RETURNING project_id
)
INSERT INTO user_projects (user_id, project_id, role)
SELECT $4, project_id, 'project manager' FROM new_project
RETURNING user_id, role, project_id
`

type CreateProjectByUserParams struct {
	ProjectID   string
	Name        string
	Description pgtype.Text
	UserID      string
}

func (q *Queries) CreateProjectByUser(ctx context.Context, arg CreateProjectByUserParams) (UserProject, error) {
	row := q.db.QueryRow(ctx, createProjectByUser,
		arg.ProjectID,
		arg.Name,
		arg.Description,
		arg.UserID,
	)
	var i UserProject
	err := row.Scan(&i.UserID, &i.Role, &i.ProjectID)
	return i, err
}

const createProjectInvitation = `-- name: CreateProjectInvitation :exec
INSERT INTO user_project_invitations (invitation_id, recipient_id, sender_id, project_id, role, status)
VALUES ($1, $2, $3, $4, $5, 0)
`

type CreateProjectInvitationParams struct {
	InvitationID string
	RecipientID  string
	SenderID     pgtype.Text
	ProjectID    string
	Role         string
}

func (q *Queries) CreateProjectInvitation(ctx context.Context, arg CreateProjectInvitationParams) error {
	_, err := q.db.Exec(ctx, createProjectInvitation,
		arg.InvitationID,
		arg.RecipientID,
		arg.SenderID,
		arg.ProjectID,
		arg.Role,
	)
	return err
}

const createTicket = `-- name: CreateTicket :one

INSERT INTO tickets (
  ticket_id,
  title,
  description,
  status,
  priority,
  created_by,
  assignee_id,
  project_id
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at
`

type CreateTicketParams struct {
	TicketID    string
	Title       string
	Description pgtype.Text
	Status      int16
	Priority    int16
	CreatedBy   string
	AssigneeID  pgtype.Text
	ProjectID   string
}

// Tickets Table Queries
func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, createTicket,
		arg.TicketID,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.Priority,
		arg.CreatedBy,
		arg.AssigneeID,
		arg.ProjectID,
	)
	var i Ticket
	err := row.Scan(
		&i.TicketID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.AssigneeID,
		&i.AssigneeUsername,
		&i.AssigneeEmail,
		&i.CreatedBy,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one


INSERT INTO users (
  user_id,
  username,
  password_hash,
  email
) VALUES ($1, $2, $3, $4)
RETURNING user_id, username, password_hash, email, created_at, updated_at
`

type CreateUserParams struct {
	UserID       string
	Username     string
	PasswordHash string
	Email        string
}

// query.sql for users, projects, and tickets tables
// Users Table Queries
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.UserID,
		arg.Username,
		arg.PasswordHash,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.PasswordHash,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM projects WHERE projects.project_id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, projectID string) error {
	_, err := q.db.Exec(ctx, deleteProject, projectID)
	return err
}

const deleteProjectInvitation = `-- name: DeleteProjectInvitation :exec
DELETE FROM user_project_invitations WHERE invitation_id = $1
`

func (q *Queries) DeleteProjectInvitation(ctx context.Context, invitationID string) error {
	_, err := q.db.Exec(ctx, deleteProjectInvitation, invitationID)
	return err
}

const deleteTicket = `-- name: DeleteTicket :exec
DELETE FROM tickets WHERE ticket_id = $1
`

func (q *Queries) DeleteTicket(ctx context.Context, ticketID string) error {
	_, err := q.db.Exec(ctx, deleteTicket, ticketID)
	return err
}

const deleteTicketsByProjectID = `-- name: DeleteTicketsByProjectID :exec
DELETE FROM tickets WHERE project_id = $1
`

func (q *Queries) DeleteTicketsByProjectID(ctx context.Context, projectID string) error {
	_, err := q.db.Exec(ctx, deleteTicketsByProjectID, projectID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID string) error {
	_, err := q.db.Exec(ctx, deleteUser, userID)
	return err
}

const getAllProjects = `-- name: GetAllProjects :many
SELECT project_id, name, description, created_by, created_at, updated_at FROM projects ORDER BY created_at DESC
`

func (q *Queries) GetAllProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.Query(ctx, getAllProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ProjectID,
			&i.Name,
			&i.Description,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTickets = `-- name: GetAllTickets :many
SELECT ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at FROM tickets ORDER BY created_at DESC
`

func (q *Queries) GetAllTickets(ctx context.Context) ([]Ticket, error) {
	rows, err := q.db.Query(ctx, getAllTickets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ticket
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.TicketID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.AssigneeID,
			&i.AssigneeUsername,
			&i.AssigneeEmail,
			&i.CreatedBy,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT user_id, username, password_hash, email, created_at, updated_at FROM users ORDER BY created_at DESC
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.PasswordHash,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAssignedTickets = `-- name: GetAssignedTickets :many
SELECT ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at FROM tickets WHERE assignee_id = $1 ORDER BY status DESC, priority ASC
`

func (q *Queries) GetAssignedTickets(ctx context.Context, assigneeID pgtype.Text) ([]Ticket, error) {
	rows, err := q.db.Query(ctx, getAssignedTickets, assigneeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ticket
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.TicketID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.AssigneeID,
			&i.AssigneeUsername,
			&i.AssigneeEmail,
			&i.CreatedBy,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectById = `-- name: GetProjectById :one
SELECT project_id, name, description, created_by, created_at, updated_at FROM projects WHERE project_id = $1
`

func (q *Queries) GetProjectById(ctx context.Context, projectID string) (Project, error) {
	row := q.db.QueryRow(ctx, getProjectById, projectID)
	var i Project
	err := row.Scan(
		&i.ProjectID,
		&i.Name,
		&i.Description,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProjectInvitationById = `-- name: GetProjectInvitationById :one
SELECT invitation_id, sender_id, recipient_id, project_id, role, status, created_at, resolved_at FROM user_project_invitations WHERE invitation_id = $1 LIMIT 1
`

func (q *Queries) GetProjectInvitationById(ctx context.Context, invitationID string) (UserProjectInvitation, error) {
	row := q.db.QueryRow(ctx, getProjectInvitationById, invitationID)
	var i UserProjectInvitation
	err := row.Scan(
		&i.InvitationID,
		&i.SenderID,
		&i.RecipientID,
		&i.ProjectID,
		&i.Role,
		&i.Status,
		&i.CreatedAt,
		&i.ResolvedAt,
	)
	return i, err
}

const getProjectInvitationsByUser = `-- name: GetProjectInvitationsByUser :many
SELECT 
  upi.invitation_id,
  upi.recipient_id,
  upi.sender_id,
  upi.project_id,
  upi.role,
  upi.status,
  p.name AS project_name,
  u.username AS sender_username,
  u.email AS sender_email
FROM
  user_project_invitations upi
JOIN
  projects p ON upi.project_id = p.project_id
JOIN
  users u ON upi.sender_id = u.user_id
WHERE
  upi.recipient_id = $1 AND upi.status = 0
`

type GetProjectInvitationsByUserRow struct {
	InvitationID   string
	RecipientID    string
	SenderID       pgtype.Text
	ProjectID      string
	Role           string
	Status         int16
	ProjectName    string
	SenderUsername string
	SenderEmail    string
}

func (q *Queries) GetProjectInvitationsByUser(ctx context.Context, recipientID string) ([]GetProjectInvitationsByUserRow, error) {
	rows, err := q.db.Query(ctx, getProjectInvitationsByUser, recipientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectInvitationsByUserRow
	for rows.Next() {
		var i GetProjectInvitationsByUserRow
		if err := rows.Scan(
			&i.InvitationID,
			&i.RecipientID,
			&i.SenderID,
			&i.ProjectID,
			&i.Role,
			&i.Status,
			&i.ProjectName,
			&i.SenderUsername,
			&i.SenderEmail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectInvitationsByUserAndProject = `-- name: GetProjectInvitationsByUserAndProject :many
SELECT invitation_id, sender_id, recipient_id, project_id, role, status, created_at, resolved_at FROM user_project_invitations WHERE recipient_id = $1 AND project_id = $2 AND status = 0
`

type GetProjectInvitationsByUserAndProjectParams struct {
	RecipientID string
	ProjectID   string
}

func (q *Queries) GetProjectInvitationsByUserAndProject(ctx context.Context, arg GetProjectInvitationsByUserAndProjectParams) ([]UserProjectInvitation, error) {
	rows, err := q.db.Query(ctx, getProjectInvitationsByUserAndProject, arg.RecipientID, arg.ProjectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserProjectInvitation
	for rows.Next() {
		var i UserProjectInvitation
		if err := rows.Scan(
			&i.InvitationID,
			&i.SenderID,
			&i.RecipientID,
			&i.ProjectID,
			&i.Role,
			&i.Status,
			&i.CreatedAt,
			&i.ResolvedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectInvitationsByUserId = `-- name: GetProjectInvitationsByUserId :many
SELECT invitation_id, sender_id, recipient_id, project_id, role, status, created_at, resolved_at FROM user_project_invitations WHERE recipient_id = $1
`

func (q *Queries) GetProjectInvitationsByUserId(ctx context.Context, recipientID string) ([]UserProjectInvitation, error) {
	rows, err := q.db.Query(ctx, getProjectInvitationsByUserId, recipientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserProjectInvitation
	for rows.Next() {
		var i UserProjectInvitation
		if err := rows.Scan(
			&i.InvitationID,
			&i.SenderID,
			&i.RecipientID,
			&i.ProjectID,
			&i.Role,
			&i.Status,
			&i.CreatedAt,
			&i.ResolvedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectMemberRelation = `-- name: GetProjectMemberRelation :one
SELECT user_id, role, project_id FROM user_projects WHERE user_id = $1 AND project_id = $2
`

type GetProjectMemberRelationParams struct {
	UserID    string
	ProjectID string
}

func (q *Queries) GetProjectMemberRelation(ctx context.Context, arg GetProjectMemberRelationParams) (UserProject, error) {
	row := q.db.QueryRow(ctx, getProjectMemberRelation, arg.UserID, arg.ProjectID)
	var i UserProject
	err := row.Scan(&i.UserID, &i.Role, &i.ProjectID)
	return i, err
}

const getProjectMembersWithRoles = `-- name: GetProjectMembersWithRoles :many
SELECT 
  u.user_id,
  u.username,
  u.email,
  up.project_id,
  up.role
FROM 
  users u
JOIN 
  user_projects up ON u.user_id = up.user_id
JOIN 
  projects p ON up.project_id = p.project_id
WHERE 
  p.project_id = $1
`

type GetProjectMembersWithRolesRow struct {
	UserID    string
	Username  string
	Email     string
	ProjectID string
	Role      string
}

func (q *Queries) GetProjectMembersWithRoles(ctx context.Context, projectID string) ([]GetProjectMembersWithRolesRow, error) {
	rows, err := q.db.Query(ctx, getProjectMembersWithRoles, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectMembersWithRolesRow
	for rows.Next() {
		var i GetProjectMembersWithRolesRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.ProjectID,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectUsers = `-- name: GetProjectUsers :many
SELECT user_id FROM user_projects WHERE project_id = $1
`

func (q *Queries) GetProjectUsers(ctx context.Context, projectID string) ([]string, error) {
	rows, err := q.db.Query(ctx, getProjectUsers, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var user_id string
		if err := rows.Scan(&user_id); err != nil {
			return nil, err
		}
		items = append(items, user_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectsByUserId = `-- name: GetProjectsByUserId :many

SELECT p.project_id, p.name, p.description, p.created_by, p.created_at, p.updated_at FROM projects p
JOIN user_projects up ON p.project_id = up.project_id
WHERE up.user_id = $1
`

// Projects Table Queries
func (q *Queries) GetProjectsByUserId(ctx context.Context, userID string) ([]Project, error) {
	rows, err := q.db.Query(ctx, getProjectsByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ProjectID,
			&i.Name,
			&i.Description,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectsByUserIdWithTicketAndMemberInfo = `-- name: GetProjectsByUserIdWithTicketAndMemberInfo :many
WITH ProjectCounts AS (
  SELECT
    p.project_id,
    COUNT(t.project_id) AS ticket_count,
    COUNT(up.project_id) AS member_count,
    SUM(CASE WHEN t.status = 1 THEN 1 ELSE 0 END) AS open_ticket_count,
    SUM(CASE WHEN t.status = 1 AND t.assignee_id = $1 THEN 1 ELSE 0 END) AS your_ticket_count
  FROM projects p
  JOIN user_projects up ON p.project_id = up.project_id
  LEFT JOIN tickets t ON p.project_id = t.project_id
  WHERE up.user_id = $1
  GROUP BY p.project_id
)
SELECT
  p.project_id, p.name, p.description, p.created_by, p.created_at, p.updated_at,
  pc.ticket_count,
  pc.member_count,
  pc.open_ticket_count,
  pc.your_ticket_count
FROM projects p
JOIN ProjectCounts pc ON p.project_id = pc.project_id
`

type GetProjectsByUserIdWithTicketAndMemberInfoRow struct {
	ProjectID       string
	Name            string
	Description     pgtype.Text
	CreatedBy       pgtype.Text
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
	TicketCount     int64
	MemberCount     int64
	OpenTicketCount int64
	YourTicketCount int64
}

// This will also have a your_tickets (assigned to you), open_tickets, and project_member_count
func (q *Queries) GetProjectsByUserIdWithTicketAndMemberInfo(ctx context.Context, assigneeID pgtype.Text) ([]GetProjectsByUserIdWithTicketAndMemberInfoRow, error) {
	rows, err := q.db.Query(ctx, getProjectsByUserIdWithTicketAndMemberInfo, assigneeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProjectsByUserIdWithTicketAndMemberInfoRow
	for rows.Next() {
		var i GetProjectsByUserIdWithTicketAndMemberInfoRow
		if err := rows.Scan(
			&i.ProjectID,
			&i.Name,
			&i.Description,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TicketCount,
			&i.MemberCount,
			&i.OpenTicketCount,
			&i.YourTicketCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTicketById = `-- name: GetTicketById :one
SELECT ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at FROM tickets WHERE ticket_id = $1
`

func (q *Queries) GetTicketById(ctx context.Context, ticketID string) (Ticket, error) {
	row := q.db.QueryRow(ctx, getTicketById, ticketID)
	var i Ticket
	err := row.Scan(
		&i.TicketID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.AssigneeID,
		&i.AssigneeUsername,
		&i.AssigneeEmail,
		&i.CreatedBy,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTicketsByProjectId = `-- name: GetTicketsByProjectId :many
SELECT ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at FROM tickets WHERE project_id = $1 ORDER BY status DESC, priority ASC
`

func (q *Queries) GetTicketsByProjectId(ctx context.Context, projectID string) ([]Ticket, error) {
	rows, err := q.db.Query(ctx, getTicketsByProjectId, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ticket
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.TicketID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.AssigneeID,
			&i.AssigneeUsername,
			&i.AssigneeEmail,
			&i.CreatedBy,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTicketsByStatus = `-- name: GetTicketsByStatus :many
SELECT ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at FROM tickets WHERE status = $1 ORDER BY created_at DESC
`

func (q *Queries) GetTicketsByStatus(ctx context.Context, status int16) ([]Ticket, error) {
	rows, err := q.db.Query(ctx, getTicketsByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ticket
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.TicketID,
			&i.Title,
			&i.Description,
			&i.Status,
			&i.Priority,
			&i.AssigneeID,
			&i.AssigneeUsername,
			&i.AssigneeEmail,
			&i.CreatedBy,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, username, password_hash, email, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.PasswordHash,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT user_id, username, password_hash, email, created_at, updated_at FROM users WHERE user_id = $1
`

func (q *Queries) GetUserById(ctx context.Context, userID string) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.PasswordHash,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserProjects = `-- name: GetUserProjects :many
SELECT project_id FROM user_projects WHERE user_id = $1
`

func (q *Queries) GetUserProjects(ctx context.Context, userID string) ([]string, error) {
	rows, err := q.db.Query(ctx, getUserProjects, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var project_id string
		if err := rows.Scan(&project_id); err != nil {
			return nil, err
		}
		items = append(items, project_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUsersWithOpenProjectInvitations = `-- name: GetUsersWithOpenProjectInvitations :many
SELECT u.user_id, u.username, u.email, upi.role 
FROM user_project_invitations upi
JOIN users u ON upi.recipient_id = u.user_id
WHERE upi.project_id = $1 AND upi.status = 0
`

type GetUsersWithOpenProjectInvitationsRow struct {
	UserID   string
	Username string
	Email    string
	Role     string
}

func (q *Queries) GetUsersWithOpenProjectInvitations(ctx context.Context, projectID string) ([]GetUsersWithOpenProjectInvitationsRow, error) {
	rows, err := q.db.Query(ctx, getUsersWithOpenProjectInvitations, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersWithOpenProjectInvitationsRow
	for rows.Next() {
		var i GetUsersWithOpenProjectInvitationsRow
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Email,
			&i.Role,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeUserFromProject = `-- name: RemoveUserFromProject :exec
DELETE FROM user_projects WHERE user_id = $1 AND project_id = $2
`

type RemoveUserFromProjectParams struct {
	UserID    string
	ProjectID string
}

func (q *Queries) RemoveUserFromProject(ctx context.Context, arg RemoveUserFromProjectParams) error {
	_, err := q.db.Exec(ctx, removeUserFromProject, arg.UserID, arg.ProjectID)
	return err
}

const removeUserFromProjectTickets = `-- name: RemoveUserFromProjectTickets :exec
UPDATE tickets SET assignee_id = NULL WHERE assignee_id = $1 AND project_id = $2
`

type RemoveUserFromProjectTicketsParams struct {
	AssigneeID pgtype.Text
	ProjectID  string
}

func (q *Queries) RemoveUserFromProjectTickets(ctx context.Context, arg RemoveUserFromProjectTicketsParams) error {
	_, err := q.db.Exec(ctx, removeUserFromProjectTickets, arg.AssigneeID, arg.ProjectID)
	return err
}

const searchUserByProject = `-- name: SearchUserByProject :many
SELECT u.user_id, u.username, u.password_hash, u.email, u.created_at, u.updated_at FROM users u
JOIN user_projects up ON u.user_id = up.user_id
WHERE up.project_id = $1
AND (u.username LIKE '%' || $2 || '%' OR u.email LIKE '%' || $2 || '%')
LIMIT 5
`

type SearchUserByProjectParams struct {
	ProjectID string
	Column2   pgtype.Text
}

func (q *Queries) SearchUserByProject(ctx context.Context, arg SearchUserByProjectParams) ([]User, error) {
	rows, err := q.db.Query(ctx, searchUserByProject, arg.ProjectID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.PasswordHash,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchUserOutsideProject = `-- name: SearchUserOutsideProject :many
SELECT u.user_id, u.username, u.password_hash, u.email, u.created_at, u.updated_at
FROM users u
LEFT JOIN user_projects up ON u.user_id = up.user_id AND up.project_id = $1
LEFT JOIN user_project_invitations upi ON u.user_id = upi.recipient_id AND upi.project_id = $1
WHERE up.user_id IS NULL
AND upi.invitation_id IS NULL
AND (u.username LIKE '%' || $2 || '%' OR u.email LIKE '%' || $2 || '%')
LIMIT 5
`

type SearchUserOutsideProjectParams struct {
	ProjectID string
	Column2   pgtype.Text
}

func (q *Queries) SearchUserOutsideProject(ctx context.Context, arg SearchUserOutsideProjectParams) ([]User, error) {
	rows, err := q.db.Query(ctx, searchUserOutsideProject, arg.ProjectID, arg.Column2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.PasswordHash,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setTicketAssignee = `-- name: SetTicketAssignee :one
UPDATE tickets 
SET 
  assignee_id = $2,
  assignee_username = u.username,
  assignee_email = u.email,
  updated_at = CURRENT_TIMESTAMP
FROM users u
WHERE 
  tickets.ticket_id = $1 AND 
  u.user_id = $2
RETURNING tickets.ticket_id, tickets.title, tickets.description, tickets.status, tickets.priority, tickets.assignee_id, tickets.assignee_username, tickets.assignee_email, tickets.created_by, tickets.project_id, tickets.created_at, tickets.updated_at
`

type SetTicketAssigneeParams struct {
	TicketID   string
	AssigneeID pgtype.Text
}

func (q *Queries) SetTicketAssignee(ctx context.Context, arg SetTicketAssigneeParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, setTicketAssignee, arg.TicketID, arg.AssigneeID)
	var i Ticket
	err := row.Scan(
		&i.TicketID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.AssigneeID,
		&i.AssigneeUsername,
		&i.AssigneeEmail,
		&i.CreatedBy,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const setTicketPrioirty = `-- name: SetTicketPrioirty :one
UPDATE tickets SET 
  priority = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at
`

type SetTicketPrioirtyParams struct {
	TicketID string
	Priority int16
}

func (q *Queries) SetTicketPrioirty(ctx context.Context, arg SetTicketPrioirtyParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, setTicketPrioirty, arg.TicketID, arg.Priority)
	var i Ticket
	err := row.Scan(
		&i.TicketID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.AssigneeID,
		&i.AssigneeUsername,
		&i.AssigneeEmail,
		&i.CreatedBy,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const setTicketStatus = `-- name: SetTicketStatus :one
UPDATE tickets SET 
  status = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at
`

type SetTicketStatusParams struct {
	TicketID string
	Status   int16
}

func (q *Queries) SetTicketStatus(ctx context.Context, arg SetTicketStatusParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, setTicketStatus, arg.TicketID, arg.Status)
	var i Ticket
	err := row.Scan(
		&i.TicketID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.AssigneeID,
		&i.AssigneeUsername,
		&i.AssigneeEmail,
		&i.CreatedBy,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateProject = `-- name: UpdateProject :one
UPDATE projects SET 
  name = $2,
  description = $3,
  updated_at = CURRENT_TIMESTAMP
WHERE project_id = $1
RETURNING project_id, name, description, created_by, created_at, updated_at
`

type UpdateProjectParams struct {
	ProjectID   string
	Name        string
	Description pgtype.Text
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (Project, error) {
	row := q.db.QueryRow(ctx, updateProject, arg.ProjectID, arg.Name, arg.Description)
	var i Project
	err := row.Scan(
		&i.ProjectID,
		&i.Name,
		&i.Description,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTicket = `-- name: UpdateTicket :one
UPDATE tickets SET 
  title = $2,
  description = $3,
  status = $4,
  priority = $5,
  assignee_id = $6,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING ticket_id, title, description, status, priority, assignee_id, assignee_username, assignee_email, created_by, project_id, created_at, updated_at
`

type UpdateTicketParams struct {
	TicketID    string
	Title       string
	Description pgtype.Text
	Status      int16
	Priority    int16
	AssigneeID  pgtype.Text
}

func (q *Queries) UpdateTicket(ctx context.Context, arg UpdateTicketParams) (Ticket, error) {
	row := q.db.QueryRow(ctx, updateTicket,
		arg.TicketID,
		arg.Title,
		arg.Description,
		arg.Status,
		arg.Priority,
		arg.AssigneeID,
	)
	var i Ticket
	err := row.Scan(
		&i.TicketID,
		&i.Title,
		&i.Description,
		&i.Status,
		&i.Priority,
		&i.AssigneeID,
		&i.AssigneeUsername,
		&i.AssigneeEmail,
		&i.CreatedBy,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET 
  username = $2,
  password_hash = $3,
  email = $4,
  updated_at = CURRENT_TIMESTAMP
WHERE user_id = $1
RETURNING user_id, username, password_hash, email, created_at, updated_at
`

type UpdateUserParams struct {
	UserID       string
	Username     string
	PasswordHash string
	Email        string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.UserID,
		arg.Username,
		arg.PasswordHash,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.PasswordHash,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
