package ports

import "github.com/dsgomes/rest-api/internal/core/domain"

type TodoUseCase interface {
	GetAll() ([]domain.Todo, error)
	Get(id string) (*domain.Todo, error)
	Insert(todo *domain.Todo) (string, error)
	Update(id string, todo *domain.Todo) (int64, error)
	Delete(id string) (int64, error)
}
