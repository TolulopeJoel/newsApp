package main

import (
	"time"

	"github.com/tolulopejoel/newsApp/internal/database"
)

type Article struct {
	ID          int32 `json:"id"`
	Title       string
	Summary     string
	Content     string
	IsPublished bool
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func databaseArticleToArticle(article database.Article) Article {
	return Article{
		ID:          article.ID,
		Title:       nullStringToString(article.Title),
		Summary:     nullStringToString(article.Summary),
		Content:     article.Content,
		IsPublished: article.IsPublished,
		PublishedAt: nullTimeToTime(article.PublishedAt),
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
}

func databaseArticlesToArticles(articles []database.Article) []Article {
	result := make([]Article, len(articles))
	for _, article := range articles {
		result = append(result, databaseArticleToArticle(article))
	}
	return result
}
