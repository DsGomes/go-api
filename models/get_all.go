package models

import (
	"github.com/dsgomes/rest-api/db"
	"github.com/dsgomes/rest-api/entities"
)

func GetAll() (todos []entities.Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM todos")
	if err != nil {
		return
	}

	for rows.Next() {
		var todo entities.Todo

		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
