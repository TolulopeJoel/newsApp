package main

import (
	"context"
	"log"
	"time"

	"github.com/tolulopejoel/newsApp/news"
)

func startBackgroundWorkers(ctx context.Context, cfg *apiConfig) {
	// scrape for new articles
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("News sources worker shutting down")
				return
			default:
				sources, err := cfg.DB.GetAllSources(ctx)
				if err != nil {
					log.Printf("Error getting news sources: %v", err)
					time.Sleep(5 * time.Minute)
					continue
				}

				formatted := news.ConvertSlice(sources, news.DatabaseSourceToSource)
				news.FetchNewsArticles(formatted)

				time.Sleep(1 * time.Hour)
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
				articles, err := cfg.DB.GetAllUnprocessedArticles(ctx)
				if err != nil {
					log.Printf("Error getting unprocessed articles: %v", err)
					time.Sleep(1 * time.Minute)
					continue
				}

				for _, article := range articles {
					news.Analyse(article)
				}

				time.Sleep(10 * time.Minute)
			}
		}
	}()

	// pusblish articles
	go func() {

	}()
}
