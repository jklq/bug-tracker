-- migrate:up

-- Create users table
CREATE TABLE users (
    user_id TEXT PRIMARY KEY NOT NULL,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create projects table
CREATE TABLE projects (
    project_id TEXT PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    created_by TEXT REFERENCES users(user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create tickets table
CREATE TABLE tickets (
    ticket_id TEXT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    status SMALLINT NOT NULL,
    priority SMALLINT NOT NULL,
    assigned_to TEXT REFERENCES users(user_id),
    created_by TEXT REFERENCES users(user_id) NOT NULL,
    project_id TEXT REFERENCES projects(project_id) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create the function to update 'updated_at' in users
CREATE OR REPLACE FUNCTION update_updated_at_column_users()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW(); 
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- Create triggers for updating 'updated_at' in users, projects, and tickets
CREATE TRIGGER update_user_modtime 
BEFORE UPDATE ON users 
FOR EACH ROW 
EXECUTE FUNCTION update_updated_at_column_users();

CREATE TRIGGER update_project_modtime 
BEFORE UPDATE ON projects 
FOR EACH ROW 
EXECUTE FUNCTION update_updated_at_column_users();

CREATE TRIGGER update_ticket_modtime 
BEFORE UPDATE ON tickets 
FOR EACH ROW 
EXECUTE FUNCTION update_updated_at_column_users();

-- migrate:down

-- Drop the triggers on users, projects, and tickets table
DROP TRIGGER IF EXISTS update_user_modtime ON users;
DROP TRIGGER IF EXISTS update_ticket_modtime ON tickets;
DROP TRIGGER IF EXISTS update_project_modtime ON projects;

-- Drop the users, projects, and tickets tables
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS users;

-- Drop the function update_updated_at_column_users
DROP FUNCTION IF EXISTS update_updated_at_column_users();
