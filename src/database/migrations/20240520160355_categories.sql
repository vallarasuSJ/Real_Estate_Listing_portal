-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories(
    id UUID PRIMARY KEY,
    "name" TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
