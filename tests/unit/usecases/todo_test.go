package usecases_test

import (
	"testing"
	"time"

	"github.com/dsgomes/rest-api/internal/core/domain"
	"github.com/dsgomes/rest-api/internal/core/usecases"
	"github.com/dsgomes/rest-api/tests/unit/mocks"
	"github.com/stretchr/testify/assert"
)

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
	mockRepo := new(mocks.MockTodoRepository)
	mockRepo.On("GetAll").Return(mockTodos, nil)
	useCase := usecases.NewTodoUseCase(mockRepo)

	todos, err := useCase.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, mockTodos, todos)

	mockRepo.AssertExpectations(t)
}

func TestGet(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	mockRepo.On("Get", "ABC").Return(mockTodos[0], nil)
	useCase := usecases.NewTodoUseCase(mockRepo)

	todos, err := useCase.Get("ABC")

	assert.NoError(t, err)
	assert.Equal(t, mockTodos[0], *todos)

	mockRepo.AssertExpectations(t)
}
