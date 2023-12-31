-- migrate:up

-- Create user_projects table
CREATE TABLE user_projects (
    user_id TEXT REFERENCES users(user_id),
    project_id TEXT REFERENCES projects(project_id),
    PRIMARY KEY (user_id, project_id)
);

-- migrate:down

-- Drop the user_projects table
DROP TABLE IF EXISTS user_projects;
