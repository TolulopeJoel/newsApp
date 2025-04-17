// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
	"time"
)

type Article struct {
	ID          int32
	Title       sql.NullString
	Summary     sql.NullString
	Content     string
	IsPublished bool
	PublishedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SourceID    int32
	ImageUrl    sql.NullString
	IsProcessed bool
	HookTitle   sql.NullString
}

type Source struct {
	ID        int32
	Name      sql.NullString
	FeedUrl   sql.NullString
	CreatedAt time.Time
}
