-- name: CreateArticle :one
INSERT INTO articles (source_id, title, summary, content) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllArticles :many
SELECT *
FROM articles
ORDER BY created_at DESC;

-- name: GetAllPublishedArticles :many
SELECT *
FROM articles
WHERE is_published = true
ORDER BY created_at DESC;

-- name: GetArticleById :one 
SELECT *
FROM articles
WHERE id = $1
LIMIT 1;

-- name: GetAllUnprocessedArticles :many
SELECT *
FROM articles
WHERE is_processed = false
ORDER BY created_at;

-- name: GetArticleBySourceIdAndTitle :one
SELECT *
FROM articles
WHERE source_id = $1
AND title = $2;

-- name: UpsertArticle :exec
INSERT INTO articles (source_id, title, content)
VALUES ($1, $2, $3)
ON CONFLICT (source_id, title) DO NOTHING;
