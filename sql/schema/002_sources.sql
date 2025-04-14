-- +goose Up
CREATE TABLE sources (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    feed_url VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE sources;