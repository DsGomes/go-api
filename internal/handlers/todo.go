package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dsgomes/rest-api/internal/core/domain"
	"github.com/dsgomes/rest-api/internal/core/ports"
	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	todoUseCase ports.TodoUseCase
}

func NewTodoHandler(todoUseCase ports.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

func (t *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	todos, err := t.todoUseCase.GetAll()
	if err != nil {
		log.Printf("[Todo] Get all error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (t *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	todo, err := t.todoUseCase.Get(id)
	if err != nil {
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

func (t *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var todo *domain.Todo
	var resp map[string]any

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("[Todo] Decode error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	id, err := t.todoUseCase.Insert(todo)

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error while trying insert todo: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo succesfully created: %s", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (t *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	var todo *domain.Todo

	id := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("[Todo] Decode error: %v", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	rows, err := t.todoUseCase.Update(string(id), todo)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if rows > 1 {
		log.Printf("[Todo] Error: updated %d registers", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Updated succesfully",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (t *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	rows, err := t.todoUseCase.Delete(id)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}

	if rows > 1 {
		log.Printf("[Todo] Error: deleted %d registers", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Deleted succesfully",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
