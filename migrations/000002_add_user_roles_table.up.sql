CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    level INTEGER NOT NULL DEFAULT 1,
    description TEXT
);

INSERT INTO roles
    (name, description, level) 
VALUES
    ('user', 'A user can create posts and comments', 1),
    ('moderator', 'A moderator can update other users posts', 2),
    ('admin', 'An admin can update and delete other users posts', 3)