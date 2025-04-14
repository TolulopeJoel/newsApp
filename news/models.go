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

type Source struct {
	ID      int32 `json:"id"`
	Name    string
	FeedUrl string
}

func DatabaseArticleToArticle(article database.Article) Article {
	return Article{
		ID:          article.ID,
		Title:       article.Title.String,
		Summary:     article.Summary.String,
		Content:     article.Content,
		IsPublished: article.IsPublished,
		PublishedAt: article.PublishedAt.Time,
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

func DatabaseSourceToSource(source database.Source) Source {
	return Source{
		ID:      source.ID,
		Name:    source.Name.String,
		FeedUrl: source.FeedUrl.String,
	}
}

func DatabaseSourcesToSources(sources []database.Source) []Source {
	result := make([]Source, len(sources))
	for _, source := range sources {
		result = append(result, DatabaseSourceToSource(source))
	}
	return result
}
