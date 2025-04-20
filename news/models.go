package news

import (
	"time"

	"github.com/tolulopejoel/newsApp/internal/database"
)

type Article struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	IsPublished bool      `json:"is_published"`
	IsProcessed bool      `json:"is_processed"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	SourceID    int32     `json:"source_id"`
	ImageUrl    string    `json:"image_url,omitempty"`
	HookTitle   string    `json:"hook_title,omitempty"`
}

type Source struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	FeedUrl   string    `json:"feed_url"`
	CreatedAt time.Time `json:"created_at"`
}

func DatabaseArticleToArticle(article database.Article) Article {
	return Article{
		ID:          article.ID,
		Title:       article.Title.String,
		Summary:     article.Summary.String,
		Content:     article.Content,
		IsPublished: article.IsPublished,
		IsProcessed: article.IsProcessed,
		PublishedAt: article.PublishedAt.Time,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
		SourceID:    article.SourceID,
		ImageUrl:    article.ImageUrl.String,
		HookTitle:   article.HookTitle.String,
	}
}

func DatabaseSourceToSource(source database.Source) Source {
	return Source{
		ID:        source.ID,
		Name:      source.Name.String,
		FeedUrl:   source.FeedUrl.String,
		CreatedAt: source.CreatedAt,
	}
}

// ConvertSlice converts a slice of database models to a slice of domain models
func ConvertSlice[T, U any](items []T, converter func(T) U) []U {
	result := make([]U, 0, len(items))
	for _, item := range items {
		result = append(result, converter(item))
	}
	return result
}
