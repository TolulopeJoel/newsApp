-- name: CreateArticle :one
INSERT INTO articles (title, summary, content) 
VALUES ($1, $2, $3)
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
