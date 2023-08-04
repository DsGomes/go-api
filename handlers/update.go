package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dsgomes/rest-api/models"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

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

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Decode error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Update error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if rows > 1 {
		log.Printf("Error: updated %d registers", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Updated succesfully",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
