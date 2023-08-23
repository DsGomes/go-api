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

func (t *todoUseCase) Get(id string) (*domain.Todo, error) {
	todo, err := t.todoRepo.Get(id)
	if err != nil {
		log.Printf("[Todo] Update error: %v", err)
		return nil, err
	}
	return &todo, nil
}

func (t *todoUseCase) Insert(todo *domain.Todo) (int64, error) {
	entity := domain.NewTodo(todo.Title, todo.Description, todo.Done)
	id, err := t.todoRepo.Insert(entity)

	if err != nil {
		log.Printf("[Todo] Insert error: %v", err)
		return 0, err
	}

	return id, nil
}

func (t *todoUseCase) Update(id string, todo *domain.Todo) (int64, error) {
	rows, err := t.todoRepo.Update(id, todo)
	if err != nil {
		log.Printf("[Todo] Update error: %v", err)
		return 0, err
	}

	return rows, nil
}

func (t *todoUseCase) Delete(id string) (int64, error) {
	rows, err := t.todoRepo.Delete(id)
	if err != nil {
		log.Printf("[Todo] Delete error: %v", err)
		return 0, err
	}

	return rows, nil
}
