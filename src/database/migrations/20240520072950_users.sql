-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY,
    email TEXT,
    gender TEXT,
    contact_number TEXT,
    role_id UUID REFERENCES roles(id),
    "password" TEXT,
    "username" TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE
);
 
 
CREATE UNIQUE INDEX users_email_role_uidx ON users using BTREE(email,"role_id");
CREATE UNIQUE INDEX users_contact_number_role_uidx ON users using BTREE(contact_number,"role_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
