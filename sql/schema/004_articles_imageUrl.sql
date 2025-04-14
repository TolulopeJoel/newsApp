-- +goose Up
ALTER TABLE articles
ADD COLUMN image_url TEXT CHECK (image_url ~ '^https?://');

-- +goose Down
ALTER TABLE articles DROP COLUMN image_url;
