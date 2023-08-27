package mocks

import (
	"github.com/dsgomes/rest-api/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) GetAll() ([]domain.Todo, error) {
	args := m.Called()
	return args.Get(0).([]domain.Todo), args.Error(1)
}

func (m *MockTodoRepository) Get(id string) (domain.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m *MockTodoRepository) Insert(todo *domain.Todo) (string, error) {
	args := m.Called(todo)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockTodoRepository) Update(id string, todo *domain.Todo) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTodoRepository) Delete(id string) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}
