package usecases_test

import (
	"testing"
	"time"

	"github.com/dsgomes/rest-api/internal/core/domain"
	"github.com/dsgomes/rest-api/internal/core/usecases"
	"github.com/stretchr/testify/assert"
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

func (m *MockTodoRepository) Insert(todo *domain.Todo) (int64, error) {
	args := m.Called(todo)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTodoRepository) Update(id string, todo *domain.Todo) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockTodoRepository) Delete(id string) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

var mockTodos = []domain.Todo{
	{
		ID:          "123456789",
		Title:       "Primeiro todo",
		Description: "Testando primeiro todo",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	},
	{
		ID:          "12343234567",
		Title:       "Segundo todo",
		Description: "Testando segundo todo",
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	},
}

func TestGetAll(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	mockRepo.On("GetAll").Return(mockTodos, nil)
	useCase := usecases.NewTodoUseCase(mockRepo)

	todos, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, mockTodos, todos)

	mockRepo.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	mockRepo := new(MockTodoRepository)
	mockRepo.On("Get", "ABC").Return(mockTodos[0], nil)
	useCase := usecases.NewTodoUseCase(mockRepo)

	todos, err := useCase.Get("ABC")

	assert.NoError(t, err)
	assert.Equal(t, mockTodos[0], *todos)

	mockRepo.AssertExpectations(t)
}
