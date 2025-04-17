-- +goose Up
ALTER TABLE articles
ADD COLUMN hook_title TEXT;

-- +goose Down
ALTER TABLE articles DROP COLUMN hook_title;
