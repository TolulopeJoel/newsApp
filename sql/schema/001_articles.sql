-- +goose Up
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    summary TEXT,
    content TEXT NOT NULL,
    is_published BOOLEAN NOT NULL DEFAULT FALSE,

    published_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE articles;
