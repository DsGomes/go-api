package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dsgomes/rest-api/internal/repositories"
)

func List(w http.ResponseWriter, r *http.Request) {
	repository := repositories.NewTodoPostgresRepository()
	todos, err := repository.GetAll()
	if err != nil {
		log.Printf("Get all error: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
