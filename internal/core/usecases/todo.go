package usecases

import (
	"log"

	"github.com/dsgomes/rest-api/internal/core/domain"
	"github.com/dsgomes/rest-api/internal/core/ports"
)

type todoUseCase struct {
	todoRepo ports.TodoRepository
}

func NewTodoUseCase(todoRepo ports.TodoRepository) ports.TodoUseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}

func (t *todoUseCase) GetAll() ([]domain.Todo, error) {
	todos, err := t.todoRepo.GetAll()
	if err != nil {
		log.Printf("[Todo] Get all error: %v", err)
	}

	return todos, nil
}

func (t *todoUseCase) Get(id int64) (*domain.Todo, error) {
	todo, err := t.todoRepo.Get(id)
	if err != nil {
		log.Printf("[Todo] Update error: %v", err)
		return nil, err
	}
	return &todo, nil
}

func (t *todoUseCase) Insert(todo *domain.Todo) (int64, error) {
	id, err := t.todoRepo.Insert(todo)

	if err != nil {
		log.Printf("[Todo] Insert error: %v", err)
		return 0, err
	}

	return id, nil
}

func (t *todoUseCase) Update(id int64, todo *domain.Todo) (int64, error) {
	rows, err := t.todoRepo.Update(id, todo)
	if err != nil {
		log.Printf("[Todo] Update error: %v", err)
		return 0, err
	}

	return rows, nil
}

func (t *todoUseCase) Delete(id int64) (int64, error) {
	rows, err := t.todoRepo.Delete(int64(id))
	if err != nil {
		log.Printf("[Todo] Delete error: %v", err)
		return 0, err
	}

	return rows, nil
}
