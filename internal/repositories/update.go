package repositories

import (
	"time"

	"github.com/dsgomes/rest-api/db"
	"github.com/dsgomes/rest-api/internal/core/domain"
)

func Update(id int64, todo domain.Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := "UPDATE todos SET title=$2, description=$3, done=$4, updated_at=$5 WHERE id=$1"

	res, err := conn.Exec(sql, id, todo.Title, todo.Description, todo.Done, time.Now())
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
