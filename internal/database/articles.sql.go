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
INSERT INTO articles (title, summary, content) 
VALUES ($1, $2, $3)
RETURNING id, title, summary, content, is_published, published_at, created_at, updated_at
`

type CreateArticleParams struct {
	Title   sql.NullString
	Summary sql.NullString
	Content string
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (Article, error) {
	row := q.db.QueryRowContext(ctx, createArticle, arg.Title, arg.Summary, arg.Content)
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
	)
	return i, err
}

const getAllArticles = `-- name: GetAllArticles :many
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at
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
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at
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
SELECT id, title, summary, content, is_published, published_at, created_at, updated_at
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
	)
	return i, err
}
