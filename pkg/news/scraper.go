package news

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/tolulopejoel/newsApp/internal/database"

	readability "github.com/go-shiori/go-readability"
	"github.com/mmcdole/gofeed"

	"github.com/tolulopejoel/newsApp/internal/database"
)

// FetchNewsArticles fetches news articles from various sources
func FetchNewsArticles(ctx context.Context, queries *database.Queries, sources []Source) {
	var wg sync.WaitGroup
	log.Println("Scraping started...")
	defer log.Println("Scraping Completed.")

	// Process each source in its own goroutine
	for _, source := range sources {
		wg.Add(1)
		go func(src string) {
			defer wg.Done()

			// Get links from RSS feed
			links, err := getLinksFromRSS(src)
			if err != nil {
				log.Printf("Error fetching article links from %s: %v", src, err)
				return
			}

			// Process each link
			var articleWg sync.WaitGroup
			for _, link := range links {
				articleWg.Add(1)
				go func(articleLink *url.URL) {
					defer articleWg.Done()

					// Get and process the article
					articlePage, err := getArticlePage(articleLink.String())
					if err != nil {
						log.Printf("Error downloading article from %s: %v", articleLink, err)
						return
					}

					article, err := extractNewsArticleInfo(articlePage, articleLink)
					if err != nil {
						log.Printf("Error extracting article info from %s: %v", articleLink, err)
						return
					}

					// Save the article to the database
					if err := saveArticleToDB(queries, source, article); err != nil {
						log.Printf("Error saving article to database: %v", err)
					}
				}(link)
			}

			// Wait for all articles from this source to finish
			articleWg.Wait()
		}(source.FeedUrl)
	}
	// Wait for all sources to complete
	wg.Wait()
}

// get articles link from news sources rss feed
func getLinksFromRSS(feedURL string) ([]*url.URL, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing feed: %w", err)
	}

	links := make([]*url.URL, 0, len(feed.Items))
	for _, item := range feed.Items {
		link, err := url.Parse(item.Link)
		if err != nil {
			log.Printf("Error parsing URL %s: %v", item.Link, err)
			continue
		}
		links = append(links, link)
	}

	return links, nil
}

// get the content on the article page
func getArticlePage(articleLink string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", articleLink, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Add headers to make the request look more like a browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error downloading article: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading article body: %w", err)
	}

	return string(body), nil
}

func extractNewsArticleInfo(articlePage string, pageURL *url.URL) (*Article, error) {
	// Parse the content using readability
	article, err := readability.FromReader(
		strings.NewReader(string(articlePage)),
		pageURL,
	)
	if err != nil {
		log.Printf("Failed to parse article content: %v", err)
		return nil, err
	}

	return &Article{
		Title:   article.Title,
		Content: article.TextContent,
	}, nil
}

// save extracted info to database
func saveArticleToDB(queries *database.Queries, source Source, article *Article) error {
	if article == nil {
		return fmt.Errorf("article cannot be nil")
	}

	return queries.UpsertArticle(context.TODO(), database.UpsertArticleParams{
		SourceID: source.ID,
		Title: sql.NullString{
			String: article.Title,
			Valid:  article.Title != "",
		},
		Content: article.Content,
	})
}
