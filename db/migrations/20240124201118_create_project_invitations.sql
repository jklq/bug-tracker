-- migrate:up

-- Create user_project_invitations table
CREATE TABLE user_project_invitations (
    invitation_id TEXT PRIMARY KEY NOT NULL,
    sender_id TEXT REFERENCES users(user_id) ON DELETE SET NULL,
    recipient_id TEXT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    project_id TEXT NOT NULL REFERENCES projects(project_id) ON DELETE CASCADE,
    role TEXT NOT NULL,
    status SMALLINT NOT NULL, -- e.g., 0 for pending, 1 for accepted, 2 for declined
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    resolved_at TIMESTAMPTZ
);

-- Create the function to update 'resolved_at'
CREATE OR REPLACE FUNCTION update_resolved_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.resolved_at = NOW(); 
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';


-- Create trigger for updating 'resolved_at' in user_project_invitations
CREATE TRIGGER update_invitation_resolved_time 
BEFORE UPDATE ON user_project_invitations 
FOR EACH ROW 
WHEN (OLD.status <> NEW.status AND NEW.status IN (1, 2))
EXECUTE FUNCTION update_resolved_at_column();

-- migrate:down


-- Drop the trigger on user_project_invitations table
DROP TRIGGER IF EXISTS update_invitation_resolved_time ON user_project_invitations;

-- Drop the user_project_invitations table
DROP TABLE IF EXISTS user_project_invitations;

-- Drop the function update_resolved_at_column
DROP FUNCTION IF EXISTS update_resolved_at_column();
