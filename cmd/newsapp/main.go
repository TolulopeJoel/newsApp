package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/tolulopejoel/newsApp/internal/api"
	"github.com/tolulopejoel/newsApp/internal/config"
	"github.com/tolulopejoel/newsApp/internal/database"
	"github.com/tolulopejoel/newsApp/internal/worker"
)

func main() {
	// Load environment variables + configs
	godotenv.Load()
	cfg := config.LoadConfig()
	db := database.GetDB()
	defer database.CloseDB()

	queries := database.New(db)
	apiCfg := &api.ApiConfig{DB: queries}

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

	// start background processes
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go worker.StartBackgroundWorkers(ctx, apiCfg.DB)

	// start server
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	log.Println("Server running on port: " + cfg.Port)
	log.Fatal(server.ListenAndServe())
}
