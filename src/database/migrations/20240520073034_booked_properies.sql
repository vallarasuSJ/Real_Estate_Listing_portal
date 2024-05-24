-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS booked_properties(
    id UUID PRIMARY KEY,
    "user_id" UUID REFERENCES users(id),
    property_id UUID REFERENCES properties(id),
    created_at TIMESTAMP WITHOUT TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
