package main

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func handleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "something went wrong")
}
