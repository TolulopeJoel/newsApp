package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
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

	router.Mount("/v1", v1)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Server running on port: " + port)
	log.Fatal(server.ListenAndServe())
}
