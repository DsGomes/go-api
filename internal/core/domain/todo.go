package domain

import "fmt"

type Todo struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func NewTodo(title string, description string, done bool) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Done:        done,
	}
}

func (t *Todo) String() string {
	return fmt.Sprintf("%s - %s", t.Title, t.Description)
}
