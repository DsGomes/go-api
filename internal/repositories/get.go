package repositories

import (
	"github.com/dsgomes/rest-api/db"
	"github.com/dsgomes/rest-api/internal/core/domain"
)

func Get(id int64) (todo domain.Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow("SELECT * FROM todos WHERE id=$1", id)

	err = row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Done,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	return
}
