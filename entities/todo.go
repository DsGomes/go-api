package entities

type Todo struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
