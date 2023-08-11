package main

import (
	"fmt"
	"net/http"

	"github.com/dsgomes/rest-api/configs"
	"github.com/dsgomes/rest-api/internal/core/usecases"
	"github.com/dsgomes/rest-api/internal/handlers"
	"github.com/dsgomes/rest-api/internal/infra/db"
	"github.com/dsgomes/rest-api/internal/infra/repositories"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	database := db.NewPostgres()
	repository := repositories.NewTodoPostgresRepository(database)
	useCase := usecases.NewTodoUseCase(repository)
	handler := handlers.NewTodoHandler(useCase)

	r := chi.NewRouter()
	r.Post("/", handler.Create)
	r.Put("/{id}", handler.Update)
	r.Delete("/{id}", handler.Delete)
	r.Get("/", handler.List)
	r.Get("/{id}", handler.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
