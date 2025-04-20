package api

import (
	"fmt"
	"net/http"

	"github.com/tolulopejoel/newsApp/internal/database"
	"github.com/tolulopejoel/newsApp/pkg/news"
)

type ApiConfig struct {
	DB *database.Queries
}

func (cfg *ApiConfig) HandlerGetNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	articles, err := cfg.DB.GetAllPublishedArticles(r.Context())
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

	respondWithJSON(w, http.StatusOK, news.ConvertSlice(articles, news.DatabaseArticleToArticle))
}

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "something went wrong")
}
