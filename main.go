package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"github.com/tolulopejoel/newsApp/internal/database"
	"github.com/tolulopejoel/newsApp/news"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		DB: database.New(db),
	}

	router := chi.NewRouter()
	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// V1 Router
	v1 := chi.NewRouter()
	v1.Get("/healthz", handleReadiness)
	v1.Get("/error", handleError)
	v1.Get("/news", apiCfg.handlerGetNews)

	router.Mount("/v1", v1)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go StartBackgroundWorkers(&apiCfg)

	log.Println("Server running on port: " + port)
	log.Fatal(server.ListenAndServe())
}

func StartBackgroundWorkers(cfg *apiConfig) {
	go func() {
		for {
			sources, err := cfg.DB.GetAllSources(context.TODO())
			if err != nil {
				log.Println("Can't get news sources from database", err)
			}

			formatted := news.DatabaseSourcesToSources(sources)
			news.FetchNewsArticles(formatted)

			time.Sleep(1 * time.Hour)
		}
	}()

	go func() {
		for {
			articles, err := cfg.DB.GetAllUnprocessedArticles(context.TODO())
			if err != nil {
				log.Println("Can't get unpublished articles from database", err)
			}

			for _, article := range articles {
				news.Analyse(article.Content)
			}

			time.Sleep(10 * time.Minute)
		}
	}()
}
