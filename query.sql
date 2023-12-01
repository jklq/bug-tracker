-- query.sql for users, projects, and tickets tables

-- Users Table Queries

-- name: CreateUser :one
INSERT INTO users (
  username,
  password_hash,
  email
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetAllUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users SET 
  username = $2,
  password_hash = $3,
  email = $4,
  updated_at = CURRENT_TIMESTAMP
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;

-- Projects Table Queries

-- name: CreateProject :one
INSERT INTO projects (
  name,
  description,
  created_by
) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProjectById :one
SELECT * FROM projects WHERE project_id = $1;

-- name: GetAllProjects :many
SELECT * FROM projects ORDER BY created_at DESC;

-- name: UpdateProject :one
UPDATE projects SET 
  name = $2,
  description = $3,
  updated_at = CURRENT_TIMESTAMP
WHERE project_id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects WHERE project_id = $1;

-- Tickets Table Queries

-- name: CreateTicket :one
INSERT INTO tickets (
  title,
  description,
  status,
  priority,
  assigned_to,
  project_id
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetTicketById :one
SELECT * FROM tickets WHERE ticket_id = $1;

-- name: GetAllTickets :many
SELECT * FROM tickets ORDER BY created_at DESC;

-- name: UpdateTicket :one
UPDATE tickets SET 
  title = $2,
  description = $3,
  status = $4,
  priority = $5,
  assigned_to = $6,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM tickets WHERE ticket_id = $1;

-- name: GetTicketsByStatus :many
SELECT * FROM tickets WHERE status = $1 ORDER BY created_at DESC;
