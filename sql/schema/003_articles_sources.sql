-- +goose Up
ALTER TABLE articles
ADD COLUMN source_id INTEGER REFERENCES sources(id);

-- First ensure there's at least one source
INSERT INTO sources (name, feed_url) 
SELECT 'News App', 'https://newsapp.com' 
WHERE NOT EXISTS (SELECT 1 FROM sources);

-- Update all null source_ids with the first available source
UPDATE articles 
SET source_id = (SELECT MIN(id) FROM sources)
WHERE source_id IS NULL;

-- Now we can safely set NOT NULL
ALTER TABLE articles ALTER COLUMN source_id SET NOT NULL;

-- +goose Down
ALTER TABLE articles DROP COLUMN source_id;