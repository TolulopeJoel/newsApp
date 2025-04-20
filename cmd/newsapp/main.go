package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/tolulopejoel/newsApp/internal/api"
	"github.com/tolulopejoel/newsApp/internal/database"
	"github.com/tolulopejoel/newsApp/internal/worker"
)

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

	apiCfg := &api.ApiConfig{
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
	v1.Get("/healthz", api.HandleReadiness)
	v1.Get("/error", api.HandleError)
	v1.Get("/news", apiCfg.HandlerGetNews)

	router.Mount("/v1", v1)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// start background processes
	go worker.StartBackgroundWorkers(ctx, apiCfg)

	// start server
	log.Println("Server running on port: " + port)
	log.Fatal(server.ListenAndServe())
}
