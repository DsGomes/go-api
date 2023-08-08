package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dsgomes/rest-api/internal/repositories"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Id parser error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	todo, err := repositories.Get(int64(id))
	if err != nil {
		log.Printf("Update error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}