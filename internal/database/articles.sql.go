// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: articles.sql

package database

import (
	"context"
	"database/sql"
)

const createArticle = `-- name: CreateArticle :one
INSERT INTO articles (source_id, title, summary, content) 
VALUES ($1, $2, $3, $4)
RETURNING id, title, summary, content, is_published, published_at, created_at, updated_at, source_id, image_url, is_processed
`

type CreateArticleParams struct {
	SourceID int32
	Title    sql.NullString
	Summary  sql.NullString
	Content  string
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (Article, error) {
	row := q.db.QueryRowContext(ctx, createArticle,
		arg.SourceID,
		arg.Title,
		arg.Summary,
		arg.Content,
	)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Summary,
		&i.Content,
		&i.IsPublished,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SourceID,
		&i.ImageUrl,
		&i.IsProcessed,
	)
	return i, err
}

const getAllArticles = `-- name: GetAllArticles :many
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at, source_id, image_url, is_processed
FROM articles
ORDER BY created_at DESC
`

func (q *Queries) GetAllArticles(ctx context.Context) ([]Article, error) {
	rows, err := q.db.QueryContext(ctx, getAllArticles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Article
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Summary,
			&i.Content,
			&i.IsPublished,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SourceID,
			&i.ImageUrl,
			&i.IsProcessed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPublishedArticles = `-- name: GetAllPublishedArticles :many
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at, source_id, image_url, is_processed
FROM articles
WHERE is_published = true
ORDER BY created_at DESC
`

func (q *Queries) GetAllPublishedArticles(ctx context.Context) ([]Article, error) {
	rows, err := q.db.QueryContext(ctx, getAllPublishedArticles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Article
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Summary,
			&i.Content,
			&i.IsPublished,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SourceID,
			&i.ImageUrl,
			&i.IsProcessed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllUnprocessedArticles = `-- name: GetAllUnprocessedArticles :many
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at, source_id, image_url, is_processed
FROM articles
WHERE is_processed = false
ORDER BY created_at
`

func (q *Queries) GetAllUnprocessedArticles(ctx context.Context) ([]Article, error) {
	rows, err := q.db.QueryContext(ctx, getAllUnprocessedArticles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Article
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Summary,
			&i.Content,
			&i.IsPublished,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SourceID,
			&i.ImageUrl,
			&i.IsProcessed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArticleById = `-- name: GetArticleById :one
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at, source_id, image_url, is_processed
FROM articles
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetArticleById(ctx context.Context, id int32) (Article, error) {
	row := q.db.QueryRowContext(ctx, getArticleById, id)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Summary,
		&i.Content,
		&i.IsPublished,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SourceID,
		&i.ImageUrl,
		&i.IsProcessed,
	)
	return i, err
}

const getArticleBySourceIdAndTitle = `-- name: GetArticleBySourceIdAndTitle :one
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at, source_id, image_url, is_processed
FROM articles
WHERE source_id = $1
AND title = $2
`

type GetArticleBySourceIdAndTitleParams struct {
	SourceID int32
	Title    sql.NullString
}

func (q *Queries) GetArticleBySourceIdAndTitle(ctx context.Context, arg GetArticleBySourceIdAndTitleParams) (Article, error) {
	row := q.db.QueryRowContext(ctx, getArticleBySourceIdAndTitle, arg.SourceID, arg.Title)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Summary,
		&i.Content,
		&i.IsPublished,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SourceID,
		&i.ImageUrl,
		&i.IsProcessed,
	)
	return i, err
}

const upsertArticle = `-- name: UpsertArticle :exec
INSERT INTO articles (source_id, title, content)
VALUES ($1, $2, $3)
ON CONFLICT (source_id, title) DO NOTHING
`

type UpsertArticleParams struct {
	SourceID int32
	Title    sql.NullString
	Content  string
}

func (q *Queries) UpsertArticle(ctx context.Context, arg UpsertArticleParams) error {
	_, err := q.db.ExecContext(ctx, upsertArticle, arg.SourceID, arg.Title, arg.Content)
	return err
}
