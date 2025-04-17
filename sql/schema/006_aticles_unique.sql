-- +goose Up
ALTER TABLE articles
ADD CONSTRAINT unique_title_per_source UNIQUE (title, source_id);

-- +goose Down
ALTER TABLE articles DROP CONSTRAINT unique_title_per_source;
