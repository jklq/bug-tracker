-- query.sql for users, projects, and tickets tables

-- Users Table Queries

-- name: CreateUser :one
INSERT INTO users (
  user_id,
  username,
  password_hash,
  email
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE user_id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

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


-- name: SearchUserByProject :many
SELECT u.* FROM users u
JOIN user_projects up ON u.user_id = up.user_id
WHERE up.project_id = $1
AND (u.username LIKE '%' || $2 || '%' OR u.email LIKE '%' || $2 || '%')
LIMIT 5;



-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1;


-- Projects Table Queries

-- name: GetProjectsByUserId :many
SELECT p.* FROM projects p
JOIN user_projects up ON p.project_id = up.project_id
WHERE up.user_id = $1;

-- name: AddUserToProject :exec
INSERT INTO user_projects (user_id, project_id) VALUES ($1, $2);

-- name: RemoveUserFromProject :exec
DELETE FROM user_projects WHERE user_id = $1 AND project_id = $2;

-- name: GetUserProjects :many
SELECT project_id FROM user_projects WHERE user_id = $1;

-- name: GetProjectUsers :many
SELECT user_id FROM user_projects WHERE project_id = $1;


-- name: CreateProject :one
INSERT INTO projects (
  project_id,
  name,
  description,
  created_by
) VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateProjectByUser :one
WITH new_project AS (
    INSERT INTO projects (project_id, name, description, created_by)
    VALUES ($1, $2, $3, $4)
    RETURNING project_id
)
INSERT INTO user_projects (user_id, project_id)
SELECT $4, project_id FROM new_project
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
DELETE FROM projects WHERE projects.project_id = $1;

-- Tickets Table Queries

-- name: CreateTicket :one
INSERT INTO tickets (
  ticket_id,
  title,
  description,
  status,
  priority,
  created_by,
  assigned_to,
  project_id
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetTicketById :one
SELECT t.*, u.username AS assignee_username, u.email AS assignee_email
FROM tickets t
LEFT JOIN users u ON t.assigned_to = u.user_id
WHERE t.ticket_id = $1;


-- name: GetTicketsByProjectId :many
SELECT t.*, u.username AS assignee_username, u.email AS assignee_email
FROM tickets t
LEFT JOIN users u ON t.assigned_to = u.user_id
WHERE t.project_id = $1
ORDER BY t.status DESC, t.priority ASC, t.created_at DESC;

-- name: GetAllTickets :many
SELECT * FROM tickets ORDER BY created_at DESC;

-- name: SetTicketStatus :one
UPDATE tickets SET 
  status = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING *;

-- name: SetTicketPrioirty :one
UPDATE tickets SET 
  priority = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING *;

-- name: SetTicketAssignee :one
UPDATE tickets SET 
  assigned_to = $2,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING *;


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

-- name: DeleteTicketsByProjectID :exec
DELETE FROM tickets WHERE project_id = $1;

-- name: DeleteTicket :exec
DELETE FROM tickets WHERE ticket_id = $1;

-- name: GetTicketsByStatus :many
SELECT * FROM tickets WHERE status = $1 ORDER BY created_at DESC;


