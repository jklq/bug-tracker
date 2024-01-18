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

-- name: GetProjectsByUserIdWithTicketAndMemberInfo :many
SELECT 
  p.project_id,
  p.name,
  p.description,
  COUNT(DISTINCT t.ticket_id) FILTER (WHERE t.status != 0) AS open_tickets_assignee_id_user,
  COUNT(DISTINCT tp.ticket_id) FILTER (WHERE tp.status != 0) AS total_open_tickets,
  COUNT(DISTINCT u.user_id) AS project_member_count
FROM 
  projects p
JOIN 
  user_projects up ON p.project_id = up.project_id
LEFT JOIN 
  tickets t ON p.project_id = t.project_id AND t.assignee_id = up.user_id
LEFT JOIN 
  tickets tp ON p.project_id = tp.project_id AND tp.status != 0
LEFT JOIN 
  users u ON up.user_id = u.user_id
WHERE 
  up.user_id = $1
GROUP BY 
  p.project_id
ORDER BY 
  p.name;

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
  assignee_id,
  project_id
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetTicketById :one
SELECT * FROM tickets WHERE ticket_id = $1;

-- name: GetAssignedTickets :many
SELECT * FROM tickets WHERE assignee_id = $1 ORDER BY status DESC, priority ASC;

-- name: GetTicketsByProjectId :many
SELECT * FROM tickets WHERE project_id = $1 ORDER BY status DESC, priority ASC;

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
RETURNING tickets.*;

-- name: UpdateTicket :one
UPDATE tickets SET 
  title = $2,
  description = $3,
  status = $4,
  priority = $5,
  assignee_id = $6,
  updated_at = CURRENT_TIMESTAMP
WHERE ticket_id = $1
RETURNING *;

-- name: DeleteTicketsByProjectID :exec
DELETE FROM tickets WHERE project_id = $1;

-- name: DeleteTicket :exec
DELETE FROM tickets WHERE ticket_id = $1;

-- name: GetTicketsByStatus :many
SELECT * FROM tickets WHERE status = $1 ORDER BY created_at DESC;


