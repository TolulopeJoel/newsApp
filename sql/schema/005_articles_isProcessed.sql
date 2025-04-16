-- +goose Up
ALTER TABLE articles
ADD COLUMN is_processed BOOLEAN NOT NULL DEFAULT false;

-- +goose Down
ALTER TABLE articles DROP COLUMN is_processed;
