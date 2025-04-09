package news

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

func DatabaseArticleToArticle(article database.Article) Article {
	return Article{
		ID:          article.ID,
		Title:       database.NullStringToString(article.Title),
		Summary:     database.NullStringToString(article.Summary),
		Content:     article.Content,
		IsPublished: article.IsPublished,
		PublishedAt: database.NullTimeToTime(article.PublishedAt),
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
}

func DatabaseArticlesToArticles(articles []database.Article) []Article {
	result := make([]Article, len(articles))
	for _, article := range articles {
		result = append(result, DatabaseArticleToArticle(article))
	}
	return result
}
