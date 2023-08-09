package ports

import "github.com/dsgomes/rest-api/internal/core/domain"

type TodoUseCase interface {
	GetAll() ([]domain.Todo, error)
	Get(id int64) (*domain.Todo, error)
	Insert(todo *domain.Todo) (int64, error)
	Update(id int64, todo *domain.Todo) (int64, error)
	Delete(id int64) (int64, error)
}
