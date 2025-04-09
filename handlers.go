package main

import (
	"fmt"
	"net/http"

	"github.com/tolulopejoel/newsApp/internal/database"
	"github.com/tolulopejoel/newsApp/news"
)

func (apiCfg *apiConfig) handlerGetNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	articles, err := apiCfg.DB.GetAllPublishedArticles(r.Context())
	if err != nil {
		respondWithError(
			w, http.StatusInternalServerError,
			fmt.Sprintf("Failed to fetch articles %s", err),
		)
		return
	}

	if articles == nil {
		articles = []database.Article{}
	}

	respondWithJSON(w, http.StatusOK, news.DatabaseArticlesToArticles(articles))
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "something went wrong")
}
