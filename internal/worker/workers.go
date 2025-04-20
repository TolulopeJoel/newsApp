package worker

import (
	"context"
	"log"
	"time"

	"github.com/tolulopejoel/newsApp/internal/database"
	"github.com/tolulopejoel/newsApp/pkg/news"
)

func StartBackgroundWorkers(ctx context.Context, queries *database.Queries) {
	// scrape for new articles
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("News sources worker shutting down")
				return
			default:
				sources, err := queries.GetAllSources(ctx)
				if err != nil {
					log.Printf("Error getting news sources: %v", err)
					time.Sleep(5 * time.Minute)
					continue
				}

				formatted := news.ConvertSlice(sources, news.DatabaseSourceToSource)
				news.FetchNewsArticles(ctx, queries, formatted)

				time.Sleep(time.Hour)
			}
		}
	}()

	// filter articles
	go func() {

	}()

	// analyse articles
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Article processing worker shutting down")
				return
			default:
				articles, err := queries.GetAllUnprocessedArticles(ctx)
				if err != nil {
					log.Printf("Error getting unprocessed articles: %v", err)
					time.Sleep(1 * time.Minute)
					continue
				}

				for _, article := range articles {
					news.Analyse(article, queries)
				}

				time.Sleep(10 * time.Minute)
			}
		}
	}()

	// pusblish articles
	go func() {

	}()
}
