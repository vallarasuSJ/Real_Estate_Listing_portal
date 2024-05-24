-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS properties(
    id UUID PRIMARY KEY,
    "name" TEXT,
    price NUMERIC,
    location TEXT,
    "user_id" UUID REFERENCES users(id),
    category_id UUID REFERENCES categories(id),
    is_approved BOOLEAN,
    is_booked BOOLEAN,
    created_at TIMESTAMP WITHOUT TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
