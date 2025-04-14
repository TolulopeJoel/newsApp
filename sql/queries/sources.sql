-- name: GetAllSources :many
SELECT *
FROM sources
ORDER BY created_at DESC;