-- migrate:up

-- Create users table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create projects table
CREATE TABLE projects (
    project_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_by INT REFERENCES users(user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create tickets table
CREATE TABLE tickets (
    ticket_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL,
    priority TEXT NOT NULL,
    assigned_to INT REFERENCES users(user_id),
    project_id INT REFERENCES projects(project_id),
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
DROP TRIGGER IF EXISTS update_project_modtime ON projects;
DROP TRIGGER IF EXISTS update_ticket_modtime ON tickets;

-- Drop the users, projects, and tickets tables
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS tickets;

-- Drop the function update_updated_at_column_users
DROP FUNCTION IF EXISTS update_updated_at_column_users();
