package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dsgomes/rest-api/internal/repositories"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewTodoPostgresRepository()
	rows, err := repository.Delete(int64(id))
	if err != nil {
		log.Printf("Delete error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if rows > 1 {
		log.Printf("Error: deleted %d registers", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Deleted succesfully",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
