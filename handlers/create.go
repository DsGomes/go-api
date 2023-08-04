package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dsgomes/rest-api/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Decode error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	id, err := models.Insert(todo)

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error while trying insert todo: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo succesfully created: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
