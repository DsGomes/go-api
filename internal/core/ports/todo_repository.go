package ports

import "github.com/dsgomes/rest-api/internal/core/domain"

type TodoRepository interface {
	GetAll() ([]domain.Todo, error)
	Get(id string) (domain.Todo, error)
	Insert(todo *domain.Todo) (int64, error)
	Update(id string, todo *domain.Todo) (int64, error)
	Delete(id string) (int64, error)
}
