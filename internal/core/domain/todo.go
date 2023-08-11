package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func NewTodo(title string, description string, done bool) *Todo {
	u := uuid.New()
	return &Todo{
		ID:          u.String(),
		Title:       title,
		Description: description,
		Done:        done,
		CreatedAt:   time.Now(),
	}
}

func (t *Todo) String() string {
	return fmt.Sprintf("%s - %s", t.Title, t.Description)
}
