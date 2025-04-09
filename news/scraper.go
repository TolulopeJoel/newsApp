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
	"time"

	readability "github.com/go-shiori/go-readability"
	"github.com/tolulopejoel/newsApp/internal/database"

	"github.com/mmcdole/gofeed"
)

func FetchNewsArticles(sources []string) {
	for _, source := range sources {
		links, err := getLinksFromRSS(source)
		if err != nil {
			log.Printf("Error fetching article links from %s: %v", source, err)
			continue
		}

		for _, link := range links {
			articlePage, err := getArticlePage(link.String())
			if err != nil {
				log.Printf("Error downloading article from %s: %v", link, err)
				continue
			}

			article, err := extractNewsArticleInfo(articlePage, link)
			if err != nil {
				log.Printf("Error extracting article info from %s: %v", link, err)
				continue
			}

			// Save the article to the database
			err = saveArticleToDB(article)
			if err != nil {
				log.Printf("Error saving article to database: %v", err)
			}
		}
	}
}

// get articles link from news sources rss feed
func getLinksFromRSS(source string) ([]*url.URL, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(source)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSS feed: %v", err)
	}

	var links []*url.URL
	for _, item := range feed.Items {
		if item.Link != "" {
			parsedURL, err := url.Parse(item.Link)
			if err != nil {
				log.Println("Skipping invalid URL:", item.Link)
				continue
			}
			links = append(links, parsedURL)
		}
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
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}

// pass article page to go-readability + extract info
func extractNewsArticleInfo(articlePage string, pageURL *url.URL) (*Article, error) {
	// Parse the content using readability
	article, err := readability.FromReader(
		strings.NewReader(string(articlePage)),
		pageURL,
	)
	if err != nil {
		log.Fatalf("Failed to parse article content: %v", err)
	}

	return &Article{
		Title:   article.Title,
		Content: article.TextContent,
	}, nil
}

// save extracted info to database
func saveArticleToDB(article *Article) error {
	if article == nil {
		return fmt.Errorf("article cannot be nil")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return fmt.Errorf("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	queries := database.New(db)

	// save data to db
	// TODO: check if article already exists in db (unique constraint)
	_, err = queries.CreateArticle(context.TODO(), database.CreateArticleParams{
		Title: sql.NullString{
			String: article.Title,
			Valid:  article.Title != "",
		},
		Content: article.Content,
	})
	if err != nil {
		return fmt.Errorf("failed to save article: %v", err)
	}

	return nil
}
