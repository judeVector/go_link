package main

import (
	"encoding/json"
	"github/judevector/go_link/internal/store"
	"net/http"
)

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var post store.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := app.store.Posts.Create(r.Context(), &post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
